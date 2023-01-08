package v1

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

type BasicAuthStrategy struct {
	compare func(username string, password string) bool
}

func (b BasicAuthStrategy) AuthFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Info("[BasicAuthStrategy] Authentication")
		auth := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)

		if len(auth) != 2 || auth[0] != "Basic" {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("not basic auth")})
			c.Abort()
			return
		}

		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)

		if len(pair) != 2 || !b.compare(pair[0], pair[1]) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Errorf("username or password not matched")})
			c.Abort()
			return
		}

		c.Set(UsernameKey, pair[0])
		c.Next()
	}
}

func NewBasicAuthStrategy(compare func(username string, password string) bool) BasicAuthStrategy {
	return BasicAuthStrategy{
		compare: compare,
	}
}
