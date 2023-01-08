package v1

import (
	ginjwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"time"
)

//const AccessKeyRef = "accessKey"
//const SecretKeyRef = "secretKey"

type JWTLoginCredential struct {
	AccessKey string `form:"accessKey" json:"accessKey" binding:"required"`
	SecretKey string `form:"secretKey" json:"secretKey" binding:"required"`
}

type JWTResponse struct {
	Code   int       `json:"code"`
	Expire time.Time `json:"expire"`
	Token  string    `json:"token"`
}

const IdentityKeyStr = "uid"

//const IsAdminStr = "is_admin"

type JWTStrategy struct {
	ginjwt.GinJWTMiddleware
}

var _ AuthStrategy = &JWTStrategy{}

func (strategy JWTStrategy) AuthFunc() gin.HandlerFunc {
	return strategy.MiddlewareFunc()
}

func MakeJWTAuthGroup(authGroup *gin.RouterGroup, ginJWT *ginjwt.GinJWTMiddleware) error {
	authStrategy := JWTStrategy{*ginJWT}
	authGroup.Use(authStrategy.AuthFunc())

	return nil
}

func CreateJWTAuthGroup(router *gin.RouterGroup, ginJWT *ginjwt.GinJWTMiddleware, relativePath string) (*gin.RouterGroup, error) {
	authStrategy := JWTStrategy{*ginJWT}
	authGroup := router.Group(relativePath)
	authGroup.Use(authStrategy.AuthFunc())

	return authGroup, nil
}

func NewJWTAuthStrategy(
	timeout time.Duration,
	authnFn func(string, string) bool,
	authzFn func(string, *gin.Context) bool,
	jwtDomain string,
	jwtSecret string,
) (*ginjwt.GinJWTMiddleware, error) {
	ginJWT, _ := ginjwt.New(&ginjwt.GinJWTMiddleware{
		Realm:            jwtDomain,
		SigningAlgorithm: "HS256",
		Key:              []byte(jwtSecret),
		Timeout:          timeout,
		MaxRefresh:       timeout,
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginValues JWTLoginCredential
			if err := c.ShouldBind(&loginValues); err != nil {
				return "", ginjwt.ErrMissingLoginValues
			}
			accessKey := loginValues.AccessKey
			secretKey := loginValues.SecretKey

			if ok := authnFn(accessKey, secretKey); ok {
				return &User{
					Identity: accessKey,
				}, nil
			}

			return nil, ginjwt.ErrFailedAuthentication
		},
		PayloadFunc: func(data interface{}) ginjwt.MapClaims {
			if v, ok := data.(*User); ok {
				return ginjwt.MapClaims{
					IdentityKeyStr: v.Identity,
				}
			}
			return ginjwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := ginjwt.ExtractClaims(c)
			return &User{
				Identity: claims[IdentityKeyStr].(string),
			}
		},
		IdentityKey: IdentityKeyStr,
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*User); ok && authzFn(v.Identity, c) {
				c.Set(UsernameKey, v.Identity)
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.IndentedJSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		SendCookie:    true,
		TimeFunc:      time.Now,
	})
	return ginJWT, nil
}

func RegisterAuthModule(
	engine *gin.Engine,
	basePath string,
	loginPath string,
	tokenRefreshPath string,
	ginJWT *ginjwt.GinJWTMiddleware,
) *gin.RouterGroup {
	group := engine.Group(basePath)
	strategy := JWTStrategy{*ginJWT}
	group.POST(loginPath, strategy.LoginHandler)
	group.POST(tokenRefreshPath, strategy.RefreshHandler)
	return group
}
