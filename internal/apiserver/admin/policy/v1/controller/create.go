package v1

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	model "ingress-authproxy/internal/apiserver/admin/policy/v1/model"
	"ingress-authproxy/internal/utils"
	"net/http"
)

func (c2 *controller) Create(c *gin.Context) {
	log.Infoln("[GINServer] policyController: create")

	var policy model.Policy

	if err := c.ShouldBindJSON(&policy); err != nil {
		log.Errorf("ErrBind: %s\n", err)
		utils.WriteResponse(c, http.StatusBadRequest, err, nil)
		return
	}

	policy.Username = c.GetString(UsernameKey)

	if err := c2.srv.NewPolicyService().Create(&policy); err != nil {
		utils.WriteResponse(c, http.StatusInternalServerError, err, nil)
		return
	}

	utils.WriteResponse(c, http.StatusOK, nil, policy)
}
