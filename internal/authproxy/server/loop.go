package server

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"ingress-auth-proxy/internal/config"
	auth "ingress-auth-proxy/pkg/auth/v1"
	ping "ingress-auth-proxy/pkg/ping/v1"
	"strconv"
	"time"
)

func installJWTAuthGroup(router *gin.Engine) {
	ginJWT, _ := auth.RegisterAuthModule(
		router,
		AuthProxyLayout.Auth.Self,
		AuthProxyLayout.Auth.Login,
		AuthProxyLayout.Auth.Refresh,
		time.Second*2*3600,
		func(username string, password string) bool {
			return true
		},
		func(username string) bool {
			return true
		},
		config.GlobalServerDesc.Opt.JWT.Secret,
		config.GlobalServerDesc.Opt.Network.Domain)

	/** 路由组 **/

	/** 如果开启认证，则创建认证路由组，否则创建普通路由组 **/
	_, _ = auth.CreateJWTAuthGroup(router, ginJWT, AuthProxyLayout.V1.Self)

}

func installPingGroup(router *gin.Engine) {
	grp := router.Group(AuthProxyLayout.Ping)
	controller := ping.NewController(nil)
	grp.GET("", controller.Info)
}

func installGinServer() *gin.Engine {
	router := gin.Default()
	installPingGroup(router)
	installJWTAuthGroup(router)
	return router
}

func RunFunc(a AuthProxyAdmin) {
	aVal := a.(*authProxyAdmin)
	log.Infoln("uuid:", aVal.desc.Opt.UUID)
	log.Infoln("domain:", aVal.desc.Opt.Network.Domain)
	log.Infoln("port:", aVal.desc.Opt.Network.Port)
	log.Infoln("interface:", aVal.desc.Opt.Network.Interface)
	log.Infoln("mysql.hostname:", aVal.desc.Opt.MySQL.Hostname)
	log.Infoln("mysql.port:", aVal.desc.Opt.MySQL.Port)
	log.Infoln("mysql.database:", aVal.desc.Opt.MySQL.Database)
	log.Debugln("jwt.secret:", aVal.desc.Opt.JWT.Secret)

	config.SetGlobalDesc(aVal.desc)
	if aVal.desc.Opt.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	ginEngine := installGinServer()

	_ = ginEngine.Run(aVal.desc.Opt.Network.Interface + ":" + strconv.Itoa(aVal.desc.Opt.Network.Port))
}
