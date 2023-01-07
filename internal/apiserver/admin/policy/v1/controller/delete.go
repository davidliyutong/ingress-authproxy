package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"ingress-authproxy/internal/utils"
	"net/http"
)

func (c2 *controller) Delete(c *gin.Context) {
	log.Infoln("[GINServer] policyController: delete")

	s := c.GetString(UsernameKey)
	if s == "" {
		utils.WriteResponse(c, http.StatusBadRequest, errors.New("empty username"), nil)
		return
	}

	if err := c2.srv.NewPolicyService().Delete(s, c.Param("name")); err != nil {
		utils.WriteResponse(c, http.StatusInternalServerError, err, nil)

		return
	}

	var msg = "deleted policy " + c.Param("name")
	utils.WriteResponse(c, http.StatusOK, nil, msg)
}
