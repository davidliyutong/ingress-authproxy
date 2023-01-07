package v1

import (
	"ingress-authproxy/internal/apiserver/admin/resource/v1/repo"
)

type Service interface {
	NewResourceService() ResourceService
}

type service struct {
	repo repo.Repo
}

var _ Service = (*service)(nil)

func NewService(repo repo.Repo) Service {
	return &service{repo}
}
