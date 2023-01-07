package v1

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"ingress-authproxy/internal/utils"
	"net/http"
)

func (c2 controller) List(c *gin.Context) {
	log.Infoln("[GinServer] userController: list")

	users, err := c2.srv.NewUserService().List()
	if err != nil {
		utils.WriteResponse(c, http.StatusInternalServerError, err, nil)

		return
	}

	utils.WriteResponse(c, http.StatusOK, nil, users)
}
