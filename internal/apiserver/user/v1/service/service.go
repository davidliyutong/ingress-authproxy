package v1

import (
	"ingress-authproxy/internal/apiserver/admin/user/v1/repo"
)

type Service interface {
	NewUserService() UserService
}

type service struct {
	repo repo.UserRepo
}

var _ Service = (*service)(nil)

func NewService(repo repo.UserRepo) Service {
	return &service{repo}
}
