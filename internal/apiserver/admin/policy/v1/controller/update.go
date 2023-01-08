package v1

import (
	"errors"
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

	policyName := c.Param("name")
	if policyName == "" {
		utils.WriteResponse(c, http.StatusBadRequest, errors.New("empty username"), nil)
		return
	}

	policy, err := c2.srv.NewPolicyService().Get(policyName)
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
