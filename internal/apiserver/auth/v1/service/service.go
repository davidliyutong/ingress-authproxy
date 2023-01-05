package v1

import (
	"ingress-auth-proxy/internal/apiserver/admin/v1/repo"
)

type Service interface {
	NewAuthService() AuthService
}

type service struct {
	repo repo.Repo
}

var _ Service = (*service)(nil)

func NewService(repo repo.Repo) Service {
	return &service{repo}
}
