package v1

import (
	model "ingress-authproxy/internal/apiserver/admin/policy/v1/model"
	"ingress-authproxy/internal/apiserver/admin/policy/v1/repo"
	authRepo "ingress-authproxy/internal/apiserver/auth/v1/repo"
)

type PolicyService interface {
	Create(policy *model.Policy) error
	Delete(policyName string) error
	Update(policy *model.Policy) error
	Get(policyName string) (*model.Policy, error)
	List() (*model.PolicyList, error)
}

type policyService struct {
	repo repo.Repo
}

func (p *policyService) Create(policy *model.Policy) error {
	err := p.repo.PolicyRepo().Create(policy)
	if err == nil {
		authRepo.Client().AuthzRepo().Trigger()
	}
	return err
}

func (p *policyService) Delete(policyName string) error {
	err := p.repo.PolicyRepo().Delete(policyName)
	if err == nil {
		authRepo.Client().AuthzRepo().Trigger()
	}
	return err
}

func (p *policyService) Update(policy *model.Policy) error {
	err := p.repo.PolicyRepo().Update(policy)
	if err == nil {
		authRepo.Client().AuthzRepo().Trigger()
	}
	return err
}

func (p *policyService) Get(policyName string) (*model.Policy, error) {
	return p.repo.PolicyRepo().Get(policyName)
}

func (p *policyService) List() (*model.PolicyList, error) {
	return p.repo.PolicyRepo().List()
}

func newPolicyService(repo repo.Repo) PolicyService {
	return &policyService{repo: repo}
}

func (s *service) NewPolicyService() PolicyService {
	return newPolicyService(s.repo)
}

var _ PolicyService = (*policyService)(nil)
