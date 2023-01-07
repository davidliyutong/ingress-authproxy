package v1

import (
	"ingress-authproxy/internal/apiserver/auth/v1/repo"
	srv "ingress-authproxy/internal/apiserver/auth/v1/service"
)

type Controller interface {
}

type controller struct {
	srv srv.Service
}

func NewController(repo repo.Repo) Controller {
	return &controller{
		srv: srv.NewService(repo),
	}
}
