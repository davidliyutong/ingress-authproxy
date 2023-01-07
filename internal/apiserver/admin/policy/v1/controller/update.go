package v1

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	model "ingress-authproxy/internal/apiserver/admin/policy/v1/model"
	"ingress-authproxy/internal/utils"
	"net/http"
)

func (c2 *controller) Update(c *gin.Context) {
	log.Infoln("[GINServer] policyController: update")

	var r model.Policy

	if err := c.ShouldBindJSON(&r); err != nil {
		utils.WriteResponse(c, http.StatusBadRequest, err, nil)
		return
	}

	username := c.GetString(UsernameKey)
	policyName := c.Param("name")

	policy, err := c2.srv.NewPolicyService().Get(username, policyName)
	if err != nil {
		utils.WriteResponse(c, http.StatusInternalServerError, err, nil)
		return
	}

	policy.AuthzPolicy = r.AuthzPolicy
	policy.Extend = r.Extend

	if err := c2.srv.NewPolicyService().Update(policy); err != nil {
		utils.WriteResponse(c, http.StatusInternalServerError, err, nil)
		return
	}

	utils.WriteResponse(c, http.StatusOK, nil, policy)
}
