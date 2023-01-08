package server

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	userRepo "ingress-authproxy/internal/apiserver/admin/user/v1/repo"
	"ingress-authproxy/internal/config"
	"os"
	"time"
)

type authProxyAdmin struct {
	name    string
	cmd     *cobra.Command
	args    []string
	desc    *config.AuthProxyDesc
	runFunc func(AuthProxyAdmin)
}

type AuthProxyAdmin interface {
	Run()
	PrepareRun() AuthProxyAdmin
}

func (a *authProxyAdmin) PrepareRun() AuthProxyAdmin {
	desc := config.NewAuthProxyDesc()
	err := desc.Parse(a.cmd)
	if err != nil {
		log.Errorln(err)
		os.Exit(1)
		return nil
	}
	desc.PostParse()
	a.desc = &desc
	a.name = "authproxy"
	a.runFunc = RunFunc
	return a
}

func (a *authProxyAdmin) Run() {
	log.Infof("run %v", a.name)
	if a.desc == nil || a.runFunc == nil {
		log.Errorln("server is not prepared")
		os.Exit(1)
		return
	}
	a.runFunc(a)
}

func NewAuthProxyAdmin(cmd *cobra.Command, args []string) AuthProxyAdmin {
	return &authProxyAdmin{
		cmd:  cmd,
		args: args,
	}
}

func adminAuth(username string, password string) bool {
	user, err := userRepo.Client().UserRepo().Get(username)
	if err != nil {
		return false
	}

	if err := user.Compare(password); err != nil {
		return false
	}

	log.Debugln("user:", user.Name, "is admin:", user.IsAdmin)
	if !(user.IsAdmin == 1) {
		return false
	}

	user.LoginedAt = time.Now()
	_ = userRepo.Client().UserRepo().Update(user)

	return true

}

func adminAuthz(u string, c *gin.Context) bool {
	return true
}

func userAuth(username string, password string) bool {
	user, err := userRepo.Client().UserRepo().Get(username)
	if err != nil {
		return false
	}

	if err := user.Compare(password); err != nil {
		return false
	}

	log.Debugln("user:", user.Name, "is admin:", user.IsAdmin)
	user.LoginedAt = time.Now()
	_ = userRepo.Client().UserRepo().Update(user)

	return true
}

func userAuthz(u string, c *gin.Context) bool {
	return true
}
