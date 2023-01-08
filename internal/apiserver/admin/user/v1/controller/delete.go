package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"ingress-authproxy/internal/utils"
	"net/http"
)

func (c2 controller) Delete(c *gin.Context) {
	log.Infoln("[GinServer] userController: delete")

	if c.Param("name") == "" {
		utils.WriteResponse(c, http.StatusBadRequest, errors.New("name is required"), nil)
		return
	}
	if err := c2.srv.NewUserService().Delete(c.Param("name")); err != nil {
		utils.WriteResponse(c, http.StatusInternalServerError, err, nil)
		return
	}

	var msg = "deleted user " + c.Param("name")
	utils.WriteResponse(c, http.StatusOK, nil, msg)
}
