package v1

import (
	ginjwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// const AccessKeyRef = "accessKey"
// const SecretKeyRef = "secretKey"

const IdentityKeyStr = "uid"

type JWTStrategy struct {
	ginjwt.GinJWTMiddleware
}

var _ AuthStrategy = &JWTStrategy{}

func NewJWTStrategy(gjwt ginjwt.GinJWTMiddleware) JWTStrategy {
	return JWTStrategy{gjwt}
}

func (strategy JWTStrategy) AuthFunc() gin.HandlerFunc {
	return strategy.MiddlewareFunc()
}
