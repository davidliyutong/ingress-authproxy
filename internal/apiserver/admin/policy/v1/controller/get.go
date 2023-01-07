package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"ingress-authproxy/internal/utils"
	"net/http"
)

func (c2 *controller) Get(c *gin.Context) {
	log.Infoln("[GINServer] policyController: get")

	s := c.GetString(UsernameKey)
	if s == "" {
		utils.WriteResponse(c, http.StatusBadRequest, errors.New("empty username"), nil)
		return
	}

	policy, err := c2.srv.NewPolicyService().Get(s, c.Param("name"))
	if err != nil {
		utils.WriteResponse(c, http.StatusInternalServerError, err, nil)

		return
	}

	utils.WriteResponse(c, http.StatusOK, nil, policy)
}
