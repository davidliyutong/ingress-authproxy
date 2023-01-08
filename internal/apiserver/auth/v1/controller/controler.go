package v1

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"ingress-authproxy/internal/apiserver/auth/v1/repo"
	srv "ingress-authproxy/internal/apiserver/auth/v1/service"
	v1 "ingress-authproxy/pkg/auth/v1"
	"net/http"
	"strconv"
	"strings"
)

type Controller interface {
	BasicAuthz(c *gin.Context)
	Update(c *gin.Context)
}

type controller struct {
	srv srv.Service
}

func (c2 controller) Update(c *gin.Context) {
	c2.srv.NewAuthzService().Update()
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func (c2 controller) BasicAuthz(c *gin.Context) {
	log.Info("[BasicAuthzStrategy] Authentication")
	auth := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)

	if len(auth) != 2 || auth[0] != "Basic" {
		//c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("not basic auth")})
		c.Header("WWW-Authenticate", "Basic realm="+strconv.Quote("Authorization Required"))
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	payload, _ := base64.StdEncoding.DecodeString(auth[1])
	pair := strings.SplitN(string(payload), ":", 2)
	resource := c.Param("resource")

	if len(pair) != 2 || !c2.srv.NewAuthzService().Authenticate(pair[0], pair[1], resource) {
		//c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("username or password not matched")})
		c.Header("WWW-Authenticate", "Basic realm="+strconv.Quote("Authorization Required"))
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Set(v1.UsernameKey, pair[0])
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func NewController(repo repo.Repo) Controller {
	return &controller{
		srv: srv.NewService(repo),
	}
}
