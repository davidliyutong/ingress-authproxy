package v1

import (
	"github.com/ory/ladon"
	"ingress-authproxy/internal/apiserver/auth/v1/repo"
	"time"
)

type AuthzService interface {
	Authenticate(username string, password string, resource string) bool
	Update()
}

type authzService struct {
	repo repo.Repo
}

func (a authzService) Update() {
	a.repo.AuthzRepo().Trigger()
}

func (a authzService) Authenticate(username string, password string, resource string) bool {
	user, err := a.repo.AuthzRepo().UserRepo().Get(username)
	if err != nil {
		return false
	}

	if err := user.Compare(password); err != nil {
		return false
	}

	go func() {
		user.LoginedAt = time.Now()
		_ = a.repo.AuthzRepo().UserRepo().Update(user)
	}()

	authRequest := ladon.Request{
		Subject:  "users:" + username,
		Action:   "get",
		Resource: "resources:ingress-auth:" + resource,
	}
	err = a.repo.AuthzRepo().Validate(&authRequest)
	if err != nil {
		return false
	}
	return true
}

func newAuthzService(repo repo.Repo) AuthzService {
	return &authzService{repo: repo}
}

func (s *service) NewAuthzService() AuthzService {
	return newAuthzService(s.repo)
}

var _ AuthzService = (*authzService)(nil)
