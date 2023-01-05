package authproxy

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
)

type authProxy struct {
	cmd  *cobra.Command
	args []string
}

type AuthProxy interface {
	Run()
}

func (a *authProxy) Run() {

	if err := a.cmd.Execute(); err != nil {
		fmt.Printf("%v %v\n", color.RedString("Error:"), err)
		os.Exit(1)
	}

}

func NewAuthProxy(cmd *cobra.Command, args []string) AuthProxy {
	return &authProxy{
		cmd:  cmd,
		args: args,
	}
}
