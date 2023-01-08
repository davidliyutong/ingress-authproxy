package server

import (
	"database/sql"
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	policyCtl "ingress-authproxy/internal/apiserver/admin/policy/v1/controller"
	policyRepo "ingress-authproxy/internal/apiserver/admin/policy/v1/repo"
	policyRepoMysql "ingress-authproxy/internal/apiserver/admin/policy/v1/repo/mysql"
	userCtl "ingress-authproxy/internal/apiserver/admin/user/v1/controller"
	userRepo "ingress-authproxy/internal/apiserver/admin/user/v1/repo"
	userRepoMysql "ingress-authproxy/internal/apiserver/admin/user/v1/repo/mysql"
	authCtl "ingress-authproxy/internal/apiserver/auth/v1/controller"
	authRepo "ingress-authproxy/internal/apiserver/auth/v1/repo"
	authRepoLadon "ingress-authproxy/internal/apiserver/auth/v1/repo/ladon"
	"ingress-authproxy/internal/config"
	"ingress-authproxy/internal/utils"
	auth "ingress-authproxy/pkg/auth/v1"
	ping "ingress-authproxy/pkg/ping/v1"
	"ingress-authproxy/pkg/version"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func installPingGroup(router *gin.Engine) {
	grp := router.Group(AuthProxyLayout.Ping)
	controller := ping.NewController(nil)
	grp.GET("", controller.Info)
}

// InstallAPIs install generic apis.
func installMiscAPIs(router *gin.Engine) {
	router.GET("/healthz", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	router.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, version.GitVersion)
	})

	router.GET(filepath.Join(AuthProxyLayout.V1.Self, "/healthz"), func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	router.GET(filepath.Join(AuthProxyLayout.V1.Self, "/version"), func(c *gin.Context) {
		c.String(http.StatusOK, version.GitVersion)
	})

}

func installV1Group(router *gin.Engine, jwt *jwt.GinJWTMiddleware) {

	/** 路由组 **/

	/** 如果开启认证，则创建认证路由组，否则创建普通路由组 **/
	_, _ = auth.CreateJWTAuthGroup(router, jwt, AuthProxyLayout.V1.Self)

	v1API := router.Group(AuthProxyLayout.V1.Self)

	adminPath, _ := filepath.Rel(AuthProxyLayout.V1.Self, AuthProxyLayout.V1.Admin.Self)
	adminGrp := v1API.Group(adminPath)
	_ = auth.MakeJWTAuthGroup(adminGrp, jwt)
	{
		userPath := AuthProxyLayout.V1.Admin.Users
		userGrp := adminGrp.Group(userPath)
		{
			userRepoClient, err := userRepoMysql.Repo(&config.GlobalServerDesc.Opt.MySQL)
			if err != nil {
				log.Fatalf("failed to create Mysql repo: %s", err.Error())
			}
			userRepo.SetClient(userRepoClient)
			userController := userCtl.NewController(userRepoClient)

			userGrp.POST("", userController.Create)
			userGrp.DELETE(":name", userController.Delete)
			userGrp.PUT(":name", userController.Update)
			userGrp.GET(":name", userController.Get)
			userGrp.GET("", userController.List)
		}
		policyPath := AuthProxyLayout.V1.Admin.Policies
		policyGrp := adminGrp.Group(policyPath)
		{
			policyRepoClient, _ := policyRepoMysql.Repo(&config.GlobalServerDesc.Opt.MySQL)
			policyRepo.SetClient(policyRepoClient)

			policyController := policyCtl.NewController(policyRepoClient)

			policyGrp.POST("", policyController.Create)
			policyGrp.DELETE(":name", policyController.Delete)
			policyGrp.PUT(":name", policyController.Update)
			policyGrp.GET(":name", policyController.Get)
			policyGrp.GET("", policyController.List)
		}

		if config.GlobalServerDesc.Opt.Debug {
			adminGrp.GET("server/option", func(c *gin.Context) {
				c.JSON(http.StatusOK, config.GlobalServerDesc.Opt)
			})
		} else {
			adminGrp.GET("server/option", func(c *gin.Context) {
				c.JSON(http.StatusForbidden, gin.H{"code": http.StatusForbidden, "message": "config disabled in non-debug mode"})
			})
		}
		adminGrp.POST("server/shutdown", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "server shutdown"})
			c.Done()
			go func() {
				time.Sleep(time.Second * 3)
				os.Exit(0)
			}()
		})
	}

	ingressAuthPath, _ := filepath.Rel(AuthProxyLayout.V1.Self, AuthProxyLayout.V1.IngressAuth.Self)
	ingressAuthGrp := v1API.Group(ingressAuthPath)
	{
		authRepoClient, _ := authRepoLadon.Repo(userRepo.Client().UserRepo(), policyRepo.Client().PolicyRepo())
		authRepo.SetClient(authRepoClient)

		authController := authCtl.NewController(authRepoClient)

		ingressAuthGrp.GET(":resource", authController.BasicAuthz)
	}
}

func installJWTGroup(router *gin.Engine) (*jwt.GinJWTMiddleware, error) {

	ginJWT, _ := auth.RegisterAuthModule(
		router,
		AuthProxyLayout.V1.JWT.Self,
		AuthProxyLayout.V1.JWT.Login,
		AuthProxyLayout.V1.JWT.Refresh,
		time.Second*2*3600,
		adminAuth,
		adminAuthz,
		config.GlobalServerDesc.Opt.JWT.Secret,
		config.GlobalServerDesc.Opt.Network.Domain)
	return ginJWT, nil

}

func installGinServer() *gin.Engine {
	router := gin.Default()
	ginJWT, _ := installJWTGroup(router)
	installPingGroup(router)
	installV1Group(router, ginJWT)
	installMiscAPIs(router)
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

	log.Infoln("version:", version.GitVersion)
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
