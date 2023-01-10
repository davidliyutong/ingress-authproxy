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
	adminUserCtl "ingress-authproxy/internal/apiserver/admin/user/v1/controller"
	"ingress-authproxy/internal/apiserver/oidc"
	userCtl "ingress-authproxy/internal/apiserver/user/v1/controller"

	adminUserRepo "ingress-authproxy/internal/apiserver/admin/user/v1/repo"
	adminUserRepoMysql "ingress-authproxy/internal/apiserver/admin/user/v1/repo/mysql"
	authCtl "ingress-authproxy/internal/apiserver/auth/v1/controller"
	authRepo "ingress-authproxy/internal/apiserver/auth/v1/repo"
	authRepoLadon "ingress-authproxy/internal/apiserver/auth/v1/repo/ladon"
	"ingress-authproxy/internal/config"
	"ingress-authproxy/internal/utils"
	auth "ingress-authproxy/pkg/auth/v1"
	"ingress-authproxy/pkg/version"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// InstallAPIs install generic apis.
func installMiscAPIs(router *gin.Engine) {
	router.GET(AuthProxyLayout.Healthz, func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	router.GET(AuthProxyLayout.Version, func(c *gin.Context) {
		c.String(http.StatusOK, version.GitVersion)
	})

	router.GET(filepath.Join(AuthProxyLayout.V1.Self, AuthProxyLayout.V1.Healthz), func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	router.GET(filepath.Join(AuthProxyLayout.V1.Self, AuthProxyLayout.V1.Version), func(c *gin.Context) {
		c.String(http.StatusOK, version.GitVersion)
	})

}

func installV1Group(router *gin.Engine) *gin.RouterGroup {
	v1API := router.Group(AuthProxyLayout.V1.Self)
	return v1API
}

func installAdminGroup(v1API *gin.RouterGroup, jwt *jwt.GinJWTMiddleware) {

	/** 路由组 **/
	adminPath, _ := filepath.Rel(AuthProxyLayout.V1.Self, AuthProxyLayout.V1.Admin.Self)
	adminGrp, _ := auth.CreateJWTAuthGroup(v1API, jwt, adminPath)
	_ = auth.MakeJWTAuthGroup(adminGrp, jwt)
	{
		userPath := AuthProxyLayout.V1.Admin.Users
		userGrp := adminGrp.Group(userPath)
		{
			userRepoClient, err := adminUserRepoMysql.Repo(&config.GlobalServerDesc.Opt.MySQL)
			if err != nil {
				log.Fatalf("failed to create Mysql repo: %s", err.Error())
			}
			adminUserRepo.SetClient(userRepoClient)
			userController := adminUserCtl.NewController(userRepoClient)

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

		serverPath, _ := filepath.Rel(AuthProxyLayout.V1.Admin.Self, AuthProxyLayout.V1.Admin.Server.Self)
		if config.GlobalServerDesc.Opt.Debug {
			adminGrp.GET(filepath.Join(serverPath, AuthProxyLayout.V1.Admin.Server.Option), func(c *gin.Context) {
				c.JSON(http.StatusOK, config.GlobalServerDesc.Opt)
			})
		} else {
			adminGrp.GET(filepath.Join(serverPath, AuthProxyLayout.V1.Admin.Server.Option), func(c *gin.Context) {
				c.JSON(http.StatusForbidden, gin.H{"code": http.StatusForbidden, "message": "config disabled in non-debug mode"})
			})
		}
		adminGrp.POST(filepath.Join(serverPath, AuthProxyLayout.V1.Admin.Server.Shutdown), func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "server shutdown"})
			c.Done()
			go func() {
				time.Sleep(time.Second * 3)
				os.Exit(0)
			}()
		})
		adminGrp.POST(filepath.Join(serverPath, AuthProxyLayout.V1.Admin.Server.Sync), func(c *gin.Context) {
			authRepo.Client().AuthzRepo().Trigger()
			c.JSON(http.StatusOK, gin.H{"message": "triggered sync"})
			c.Done()
		})
	}
}

func installAuthGroup(v1API *gin.RouterGroup) {

	ingressAuthPath, _ := filepath.Rel(AuthProxyLayout.V1.Self, AuthProxyLayout.V1.IngressAuth.Self)
	ingressAuthGrp := v1API.Group(ingressAuthPath)
	{
		authRepoClient, _ := authRepoLadon.Repo(adminUserRepo.Client().UserRepo(), policyRepo.Client().PolicyRepo())
		authRepo.SetClient(authRepoClient)

		authController := authCtl.NewController(authRepoClient)

		ingressAuthGrp.GET(":resource", authController.BasicAuthz)
	}
}
func installUserGroup(v1API *gin.RouterGroup, strategy auth.BasicAuthStrategy) {

	userPath := AuthProxyLayout.V1.User
	userPathGrp := v1API.Group(userPath, strategy.AuthFunc())
	{
		userRepoClient := adminUserRepo.Client().UserRepo()
		userController := userCtl.NewController(userRepoClient)

		userPathGrp.GET(":name", userController.Get)
		userPathGrp.PUT(":name", userController.Update)

	}
}

func installGinServer() *gin.Engine {
	router := gin.Default()

	adminJWT, _ := auth.NewJWTAuthStrategy(time.Second*2*3600,
		adminAuth,
		adminAuthz,
		config.GlobalServerDesc.Opt.JWT.Secret,
		config.GlobalServerDesc.Opt.Network.Domain)
	_ = auth.RegisterAuthModule(router,
		AuthProxyLayout.V1.JWT.Self,
		AuthProxyLayout.V1.JWT.Login,
		AuthProxyLayout.V1.JWT.Refresh, adminJWT)

	v1API := installV1Group(router)
	installAdminGroup(v1API, adminJWT)
	installAuthGroup(v1API)

	userAuth := auth.NewBasicAuthStrategy(userAuth)
	installUserGroup(v1API, userAuth)

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

	go oidc.RunOIDCServer(&aVal.desc.Opt.OIDC)

	_ = ginEngine.Run(aVal.desc.Opt.Network.Interface + ":" + strconv.Itoa(aVal.desc.Opt.Network.Port))
}
