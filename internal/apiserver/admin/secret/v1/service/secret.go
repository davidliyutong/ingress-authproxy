package v1

import (
	"ingress-auth-proxy/internal/apiserver/admin/secret/v1/repo"
)

type SecretService interface {
}

type secretService struct {
	repo repo.Repo
}

func newAdminService(repo repo.Repo) SecretService {
	return &secretService{repo: repo}
}

func (s *service) NewSecretService() SecretService {
	return newAdminService(s.repo)
}

var _ SecretService = (*secretService)(nil)
