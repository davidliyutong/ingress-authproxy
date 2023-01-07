package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"ingress-authproxy/internal/utils"
	"net/http"
)

func (c2 controller) Get(c *gin.Context) {
	log.Infoln("[GinServer] userController: get")

	if c.Param("name") == "" {
		utils.WriteResponse(c, http.StatusBadRequest, errors.New("name is required"), nil)
	}
	user, err := c2.srv.NewUserService().Get(c.Param("name"))
	if err != nil {
		utils.WriteResponse(c, http.StatusInternalServerError, err, nil)

		return
	}

	utils.WriteResponse(c, http.StatusOK, nil, user)
}
