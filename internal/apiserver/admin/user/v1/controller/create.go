package v1

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	model "ingress-auth-proxy/internal/apiserver/admin/user/v1/model"
	"ingress-auth-proxy/internal/utils"
	"net/http"
)

func (c2 controller) Create(c *gin.Context) {
	log.Infoln("[GinServer] userController: create")

	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		log.Errorf("ErrBind: %s\n", err)
		utils.WriteResponse(c, http.StatusBadRequest, err, nil)
		return
	}

	if err := c2.srv.NewUserService().Create(&user); err != nil {
		utils.WriteResponse(c, http.StatusInternalServerError, err, nil)
		return
	}

	utils.WriteResponse(c, http.StatusOK, nil, user)
}
