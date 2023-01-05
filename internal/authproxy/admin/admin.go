package admin

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"ingress-auth-proxy/internal/config"
	"os"
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
