package v1

import (
	"ingress-auth-proxy/internal/apiserver/admin/v1/repo"
	srv "ingress-auth-proxy/internal/apiserver/admin/v1/service"
)

type Controller interface {
}

type controller struct {
	srv srv.Service
}

func NewController(repo repo.Repo, err error) Controller {
	return &controller{
		srv: srv.NewService(repo),
	}
}
