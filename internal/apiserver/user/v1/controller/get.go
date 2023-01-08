package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"ingress-authproxy/internal/utils"
	auth "ingress-authproxy/pkg/auth/v1"
	"net/http"
)

func (c2 controller) Get(c *gin.Context) {
	log.Infoln("[GinServer] userController: get")

	var name string

	if name = c.Param("name"); name == "" {
		utils.WriteResponse(c, http.StatusBadRequest, errors.New("name is required"), nil)
		return
	}

	id := c.GetString(auth.UsernameKey)
	if id != name {
		utils.WriteResponse(c, http.StatusForbidden, errors.New("not allowed"), nil)
		return
	}

	user, err := c2.srv.NewUserService().Get(name)
	if err != nil {
		utils.WriteResponse(c, http.StatusInternalServerError, err, nil)
		return
	}

	utils.WriteResponse(c, http.StatusOK, nil, user)
}
