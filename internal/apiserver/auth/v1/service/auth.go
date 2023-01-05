package v1

import (
	"ingress-auth-proxy/internal/apiserver/admin/v1/repo"
)

type AuthService interface {
}

type authService struct {
	repo repo.Repo
}

func newAuthService(repo repo.Repo) AuthService {
	return &authService{repo: repo}
}

func (s *service) NewAuthService() AuthService {
	return newAuthService(s.repo)
}

var _ AuthService = (*authService)(nil)
