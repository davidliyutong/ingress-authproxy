package v1

import (
	"github.com/gin-gonic/gin"
	"ingress-authproxy/internal/apiserver/admin/user/v1/repo"
	srv "ingress-authproxy/internal/apiserver/user/v1/service"
)

type Controller interface {
	Update(c *gin.Context)
	Get(c *gin.Context)
}

type controller struct {
	srv srv.Service
}

func NewController(repo repo.UserRepo) Controller {
	return &controller{
		srv: srv.NewService(repo),
	}
}
