package config

import (
	"github.com/spf13/cobra"
	"os"
	"path"
)

var userHomeDir, _ = os.UserHomeDir()

const DefaultPort = 50032
const DefaultInterface = "0.0.0.0"
const DefaultAppName = "authproxy"
const DefaultConfigName = "nameserver"

var DefaultConfig = path.Join(userHomeDir, ".config/"+DefaultAppName+"/"+DefaultConfigName+".yaml")

func InitCfg(cmd *cobra.Command, args []string) {

}

func Warm(cmd *cobra.Command, args []string) {

}
