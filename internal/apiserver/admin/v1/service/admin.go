package v1

import (
	"ingress-auth-proxy/internal/apiserver/admin/v1/repo"
)

type AdminService interface {
}

type adminService struct {
	repo repo.Repo
}

func newAdminService(repo repo.Repo) AdminService {
	return &adminService{repo: repo}
}

func (s *service) NewAdminService() AdminService {
	return newAdminService(s.repo)
}

var _ AdminService = (*adminService)(nil)
