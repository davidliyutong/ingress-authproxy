package v1

import (
	model "ingress-authproxy/internal/apiserver/admin/policy/v1/model"
	"ingress-authproxy/internal/apiserver/admin/policy/v1/repo"
)

type PolicyService interface {
	Create(policy *model.Policy) error
	Delete(username, policyName string) error
	Update(policy *model.Policy) error
	Get(username, policyName string) (*model.Policy, error)
	List(username string) (*model.PolicyList, error)
}

type policyService struct {
	repo repo.Repo
}

func (p *policyService) Create(policy *model.Policy) error {
	return p.repo.PolicyRepo().Create(policy)
}

func (p *policyService) Delete(username, policyName string) error {
	return p.repo.PolicyRepo().Delete(username, policyName)
}

func (p *policyService) Update(policy *model.Policy) error {
	return p.repo.PolicyRepo().Update(policy)
}

func (p *policyService) Get(username, policyName string) (*model.Policy, error) {
	return p.repo.PolicyRepo().Get(username, policyName)
}

func (p *policyService) List(username string) (*model.PolicyList, error) {
	return p.repo.PolicyRepo().List(username)
}

func newPolicyService(repo repo.Repo) PolicyService {
	return &policyService{repo: repo}
}

func (s *service) NewPolicyService() PolicyService {
	return newPolicyService(s.repo)
}

var _ PolicyService = (*policyService)(nil)
