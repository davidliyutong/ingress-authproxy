package server

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"ingress-auth-proxy/internal/config"
	"ingress-auth-proxy/internal/utils"
	auth "ingress-auth-proxy/pkg/auth/v1"
	ping "ingress-auth-proxy/pkg/ping/v1"
	"os"
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

func mustVerifyDatabase(opt config.MySQLOpt, username string, password string) error {
	dsn := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&multiStatements=true&loc=%s`,
		opt.Username,
		opt.Password,
		opt.Hostname,
		opt.Database,
		utils.GetMySQLTZFromEnv())

	db, err := sql.Open("mysql",
		dsn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
	if !config.Verify(db) {

		log.Infoln("tables not found, creating tables...")
		err := config.Warm(db)
		if err != nil {
			log.Errorln("init database failed:", err)
			return err
		}
		success := config.Verify(db)
		if !success {
			log.Errorln("previous database initialization failed")
			return err
		}
		err = config.CreateDefaultAdminUser(db, username, password)
		if err != nil {
			log.Errorln("create default admin user failed:", err)
			return err
		}

		return nil
	}
	return nil
}

func RunFunc(a AuthProxyAdmin) {
	aVal := a.(*authProxyAdmin)

	// init database
	err := mustVerifyDatabase(aVal.desc.Opt.MySQL, aVal.desc.Opt.Init.Username, aVal.desc.Opt.Init.Password)
	if err != nil {
		os.Exit(1)
	}
	aVal.desc.Opt.MySQL.Initialized = true
	if aVal.desc.Viper.ConfigFileUsed() != "" {
		utils.DumpOption(aVal.desc.Opt, aVal.desc.Viper.ConfigFileUsed(), true)
	}

	// configure gin
	config.SetGlobalDesc(aVal.desc)
	if aVal.desc.Opt.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	ginEngine := installGinServer()

	log.Infoln("uuid:", aVal.desc.Opt.UUID)
	log.Infoln("domain:", aVal.desc.Opt.Network.Domain)
	log.Infoln("port:", aVal.desc.Opt.Network.Port)
	log.Infoln("interface:", aVal.desc.Opt.Network.Interface)
	log.Infoln("mysql.initialized:", aVal.desc.Opt.MySQL.Initialized)
	log.Infoln("mysql.hostname:", aVal.desc.Opt.MySQL.Hostname)
	log.Infoln("mysql.port:", aVal.desc.Opt.MySQL.Port)
	log.Infoln("mysql.database:", aVal.desc.Opt.MySQL.Database)
	log.Debugln("jwt.secret:", aVal.desc.Opt.JWT.Secret)

	_ = ginEngine.Run(aVal.desc.Opt.Network.Interface + ":" + strconv.Itoa(aVal.desc.Opt.Network.Port))
}
