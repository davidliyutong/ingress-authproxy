package v1

import (
	"ingress-auth-proxy/internal/apiserver/admin/user/v1/repo"
)

type UserService interface {
}

type userService struct {
	repo repo.Repo
}

func newAdminService(repo repo.Repo) UserService {
	return &userService{repo: repo}
}

func (s *service) NewAdminService() UserService {
	return newAdminService(s.repo)
}

var _ UserService = (*userService)(nil)
