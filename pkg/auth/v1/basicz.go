package v1

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
)

type BasicAuthzStrategy struct {
	compare func(username string, password string, resource string) bool
	paramFn func(c *gin.Context) string
}

func (b BasicAuthzStrategy) AuthFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Info("[BasicAuthzStrategy] Authentication")
		auth := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)

		if len(auth) != 2 || auth[0] != "Basic" {
			//c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("not basic auth")})
			c.Header("WWW-Authenticate", "Basic realm="+strconv.Quote("Authorization Required"))
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)
		resource := b.paramFn(c)

		if len(pair) != 2 || !b.compare(pair[0], pair[1], resource) {
			//c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("username or password not matched")})
			c.Header("WWW-Authenticate", "Basic realm="+strconv.Quote("Authorization Required"))
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set(UsernameKey, pair[0])
		c.Next()
	}
}

func NewBasicAuthzStrategy(compare func(username string, password string, resource string) bool,
	paramFn func(c *gin.Context) string) BasicAuthzStrategy {
	return BasicAuthzStrategy{
		compare: compare,
		paramFn: paramFn,
	}
}
