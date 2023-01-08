package ladon

import (
	"context"
	"github.com/ory/ladon"
	"github.com/ory/ladon/manager/memory"
	log "github.com/sirupsen/logrus"
	policyRepo "ingress-authproxy/internal/apiserver/admin/policy/v1/repo"
	userRepo "ingress-authproxy/internal/apiserver/admin/user/v1/repo"
	repoInterface "ingress-authproxy/internal/apiserver/auth/v1/repo"
	"sync"
	"time"
)

type authzRepo struct {
	userRepoClient   userRepo.UserRepo
	policyRepoClient policyRepo.PolicyRepo
	policiesCache    []*ladon.DefaultPolicy
	warden           *ladon.Ladon
	mutex            *sync.RWMutex

	channel   chan int
	stopCtx   context.Context
	stopFn    func()
	taskGroup *sync.WaitGroup
}

func (a *authzRepo) UserRepo() userRepo.UserRepo {
	return a.userRepoClient
}

func (a *authzRepo) PolicyRepo() policyRepo.PolicyRepo {
	return a.policyRepoClient
}

func (a *authzRepo) List() ([]*ladon.DefaultPolicy, error) {
	a.mutex.RLock()
	defer a.mutex.RUnlock()
	return a.policiesCache, nil
}

type AuthzRepo interface {
	List() ([]*ladon.DefaultPolicy, error)
	Start()
	Stop()
	Trigger()
	Validate(request *ladon.Request) error
	UserRepo() userRepo.UserRepo
	PolicyRepo() policyRepo.PolicyRepo
}

var _ repoInterface.AuthzRepo = (*authzRepo)(nil)

var (
	repoInstance *authzRepo
	onceCache    sync.Once
)

func (a *authzRepo) resetPolicies(policies []*ladon.DefaultPolicy) {

	newManager := memory.NewMemoryManager()
	for _, policy := range policies {
		_ = newManager.Create(policy)
	}

	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.policiesCache = policies
	a.warden.Manager = newManager
	log.Debugf("[Authorizer] updated cache, got %d policies", len(repoInstance.policiesCache))

}

func (a *authzRepo) update() {

	policyList, _ := a.policyRepoClient.List()
	newCache := make([]*ladon.DefaultPolicy, 0, len(a.policiesCache))
	for _, v := range policyList.Items {
		newCache = append(newCache, &v.AuthzPolicy.DefaultPolicy)
	}

	a.resetPolicies(newCache)
	a.Done()
}

func (a *authzRepo) Start() {
	a.taskGroup.Wait()
	a.stopCtx, a.stopFn = context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-a.stopCtx.Done():
				a.taskGroup.Wait()
				return
			default:
				_, ok := <-a.channel
				if ok {
					time.Sleep(time.Second * 1) // Delay to avoid multiple updates
					go a.update()
				}
			}
		}
	}()

	go func() {
		time.Sleep(time.Second * 3600)
		a.Trigger()
	}()
}

func (a *authzRepo) Stop() {
	time.Sleep(time.Second * 3)
	a.stopFn()
	a.taskGroup.Wait()
}

func (a *authzRepo) Trigger() {
	a.channel <- 1
	a.taskGroup.Add(1)
}

func (a *authzRepo) Done() {
	a.taskGroup.Done()
}

func (a *authzRepo) Wait() {
	a.taskGroup.Wait()
}

func (a *authzRepo) Validate(request *ladon.Request) error {
	err := a.warden.IsAllowed(request)
	return err
}

// newAuthzRepo creates and returns a user storage.
func newAuthzRepo(userRepoClient userRepo.UserRepo, policyRepoClient policyRepo.PolicyRepo) repoInterface.AuthzRepo {
	onceCache.Do(func() {
		repoInstance = &authzRepo{
			userRepoClient:   userRepoClient,
			policyRepoClient: policyRepoClient,
			warden: &ladon.Ladon{
				Manager:     memory.NewMemoryManager(),
				AuditLogger: &ladon.AuditLoggerInfo{},
			},

			policiesCache: make([]*ladon.DefaultPolicy, 0),
			mutex:         &sync.RWMutex{},

			channel:   make(chan int, 16),
			taskGroup: &sync.WaitGroup{},
		}
		log.Infoln("[Authorizer] updating cache")
		go func() {
			time.Sleep(1 * time.Second)
			repoInstance.Start()
			repoInstance.Trigger()
			repoInstance.Wait()
			log.Infof("[Authorizer] updated cache, got %d policies", len(repoInstance.policiesCache))
		}()

	})
	return repoInstance
}
