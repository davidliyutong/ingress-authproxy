package v1

import (
	"ingress-auth-proxy/internal/apiserver/auth/v1/repo"
)

type AuthzService interface {
}

type authzService struct {
	repo repo.Repo
}

func newAuthzService(repo repo.Repo) AuthzService {
	return &authzService{repo: repo}
}

func (s *service) NewAuthzService() AuthzService {
	return newAuthzService(s.repo)
}

var _ AuthzService = (*authzService)(nil)
