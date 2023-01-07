package v1

import (
	"github.com/gin-gonic/gin"
	"ingress-authproxy/internal/apiserver/admin/policy/v1/repo"
	srv "ingress-authproxy/internal/apiserver/admin/policy/v1/service"
)

const UsernameKey = "username"

type Controller interface {
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	Get(c *gin.Context)
	List(c *gin.Context)
}

type controller struct {
	srv srv.Service
}

func NewController(repo repo.Repo) Controller {
	return &controller{
		srv: srv.NewService(repo),
	}
}
