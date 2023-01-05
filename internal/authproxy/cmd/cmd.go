package cmd

import (
	"github.com/spf13/cobra"
	"ingress-auth-proxy/internal/authproxy/server"
	"ingress-auth-proxy/internal/config"
)

var rootCmd = &cobra.Command{
	Use:   "authproxy",
	Short: "authproxy, add authentication for your applications",
	Long:  "authproxy, add authentication for your applications",
}

var serveCmd = &cobra.Command{
	Use: "serve",
	SuggestFor: []string{
		"ru", "ser",
	},
	Short: "serve start the authproxy using predefined configs.",
	Long: `serve start the authproxy using predefined configs, by the following order:

1. path specified in --config flag
2. path defined DFS_CONFIG environment variable
3. default location $HOME/.config/authproxy/nameserver.yaml, /etc/authproxy/nameserver.yaml, current directory

The parameters in the configuration file will be overwritten by the following order:

1. command line arguments
2. environment variables
`,
	Example: `  authproxy serve --config=path_to_config`,
	Run: func(cmd *cobra.Command, args []string) {
		server.NewAuthProxyAdmin(cmd, args).PrepareRun().Run()
	},
}

var initCmd = &cobra.Command{
	Use: "init",
	SuggestFor: []string{
		"ini", "in",
	},
	Short: "init create a configuration template",
	Long: `init create a configuration template. This will generate uuids, default secrets and etc. 

The configuration file can be used to launch the nameserver.
If --print flag is present, the configuration will be printed to stdout.
If --output / -o flag is present, the configuration will be saved to the path specified
Otherwise init will output configuration file to $HOME/.config/go-dfs-server/nameserver.yaml
If --yes / -y flag is present, the configuration will be overwrite without confirmation
`,
	Example: `  authproxy init --print
  authproxy init --output /path/to/authproxy.yaml
  authproxy init -o /path/to/authproxy.yaml -y`,
	Run: config.InitCfg,
}

func getRootCmd() *cobra.Command {

	serveCmd.Flags().String("config", "", "default configuration path")
	serveCmd.Flags().Int64P("port", "p", config.DefaultPort, "port that nameserver listen on")
	serveCmd.Flags().StringP("interface", "i", config.DefaultInterface, "interface that nameserver listen on, default to 0.0.0.0")

	serveCmd.Flags().String("accessKey", "", "server access key")
	serveCmd.Flags().String("secretKey", "", "server secret key")
	serveCmd.Flags().Bool("debug", false, "toggle debug logging")
	rootCmd.AddCommand(serveCmd)

	initCmd.Flags().Bool("print", false, "print config to stdout")
	initCmd.Flags().BoolP("yes", "y", false, "overwrite")
	initCmd.Flags().StringP("output", "o", config.DefaultConfig, "specify output directory")
	rootCmd.AddCommand(initCmd)

	return rootCmd
}

func Execute() {
	rootCmd := getRootCmd()
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
