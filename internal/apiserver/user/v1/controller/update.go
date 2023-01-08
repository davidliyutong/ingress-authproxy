package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	model "ingress-authproxy/internal/apiserver/admin/user/v1/model"
	"ingress-authproxy/internal/utils"
	auth "ingress-authproxy/pkg/auth/v1"
	"net/http"
)

func (c2 controller) Update(c *gin.Context) {
	log.Infoln("[GinServer] userController: update")

	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		log.Errorf("ErrBind: %s\n", err)
		utils.WriteResponse(c, http.StatusBadRequest, err, nil)
		return
	}

	id := c.GetString(auth.UsernameKey)
	if id != c.Param("name") {
		utils.WriteResponse(c, http.StatusForbidden, errors.New("not allowed"), nil)
		return
	}

	user.Name = c.Param("name")

	if err := c2.srv.NewUserService().Update(&user); err != nil {
		utils.WriteResponse(c, http.StatusInternalServerError, err, nil)
		return
	}

	utils.WriteResponse(c, http.StatusOK, nil, user)
}
