package v1

import (
	"ingress-auth-proxy/internal/apiserver/admin/policy/v1/repo"
)

type PolicyService interface {
}

type policyService struct {
	repo repo.Repo
}

func newAdminService(repo repo.Repo) PolicyService {
	return &policyService{repo: repo}
}

func (s *service) NewPolicyService() PolicyService {
	return newAdminService(s.repo)
}

var _ PolicyService = (*policyService)(nil)
