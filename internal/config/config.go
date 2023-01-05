package config

import (
	"bufio"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"ingress-auth-proxy/internal/utils"
	"os"
	"path"
	"strings"
)

var userHomeDir, _ = os.UserHomeDir()

const DefaultPort = 50032
const DefaultInterface = "0.0.0.0"
const DefaultAppName = "authproxy"
const DefaultConfigName = "nameserver"

var DefaultConfig = path.Join(userHomeDir, ".config/"+DefaultAppName+"/"+DefaultConfigName+".yaml")

const DefaultConfigSearchPath0 = "/etc/authproxy"
const DefaultConfigSearchPath1 = "./"
const DefaultConfigSearchPath2 = "/config"

type MySQLOpt struct {
	Hostname string `yaml:"hostname"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type NetworkOpt struct {
	Port      int    `yaml:"port"`
	Interface string `yaml:"interface"`
}

type LogOpt struct {
	Level string `yaml:"port"`
	Path  string `yaml:"path"`
}

type AuthProxyOpt struct {
	UUID    string     `yaml:"uuid"`
	Network NetworkOpt `yaml:"network"`
	Debug   bool       `yaml:"debug"`
	Log     LogOpt     `yaml:"log"`
	MySQL   MySQLOpt   `yaml:"mysql"`
}

type AuthProxyDesc struct {
	Opt   AuthProxyOpt
	Viper *viper.Viper
	RunID string
}

func NewAuthProxyDesc() AuthProxyDesc {
	return AuthProxyDesc{
		Opt:   NewAuthProxyOpt(),
		Viper: nil,
		RunID: "",
	}
}

func NewAuthProxyOpt() AuthProxyOpt {
	return AuthProxyOpt{
		UUID: "",
		Network: NetworkOpt{
			Port:      DefaultPort,
			Interface: DefaultInterface,
		},
		Debug: false,
		Log: LogOpt{
			Level: "info",
			Path:  "",
		},
		MySQL: MySQLOpt{
			Hostname: "",
			Port:     3306,
			Database: "authproxy",
			Username: "authproxy",
			Password: "authproxy",
		},
	}
}

func (o *AuthProxyDesc) Parse(cmd *cobra.Command) error {
	vipCfg := viper.New()
	vipCfg.SetDefault("network.port", DefaultPort)
	vipCfg.SetDefault("network.interface", DefaultInterface)
	vipCfg.SetDefault("debug", false)
	vipCfg.SetDefault("log.debug", "info")
	vipCfg.SetDefault("log.path", "./authproxy.log")
	vipCfg.SetDefault("mysql.hostname", "127.0.0.1")
	vipCfg.SetDefault("mysql.port", 3306)
	vipCfg.SetDefault("mysql.database", "authproxy")
	vipCfg.SetDefault("mysql.username", "authproxy")
	vipCfg.SetDefault("mysql.password", "authproxy")

	if configFileCmd, err := cmd.Flags().GetString("config"); err == nil && configFileCmd != "" {
		vipCfg.SetConfigFile(configFileCmd)
	} else {
		configFileEnv := os.Getenv("AUTHPROXY_CONFIG")
		if configFileEnv != "" {
			vipCfg.SetConfigFile(configFileEnv)
		} else {
			vipCfg.SetConfigName(DefaultConfigName)
			vipCfg.SetConfigType("yaml")
			vipCfg.AddConfigPath(DefaultConfigSearchPath0)
			vipCfg.AddConfigPath(DefaultConfigSearchPath1)
			vipCfg.AddConfigPath(DefaultConfigSearchPath2)
		}
	}
	vipCfg.WatchConfig()

	vipCfg.SetEnvPrefix("AUTHPROXY")
	vipCfg.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	vipCfg.AutomaticEnv()

	_ = vipCfg.BindPFlag("uuid", cmd.Flags().Lookup("uuid"))
	_ = vipCfg.BindPFlag("network.port", cmd.Flags().Lookup("port"))
	_ = vipCfg.BindPFlag("network.interface", cmd.Flags().Lookup("interface"))
	_ = vipCfg.BindPFlag("debug", cmd.Flags().Lookup("debug"))

	// If a config file is found, read it in.
	if err := vipCfg.ReadInConfig(); err == nil {
		log.Debugln("using config file:", vipCfg.ConfigFileUsed())
	} else {
		log.Warnln(err)
	}

	if err := vipCfg.Unmarshal(&o.Opt); err != nil {
		log.Fatalln("failed to unmarshal config")
		os.Exit(1)
	}
	o.Viper = vipCfg
	return nil
}

func (o *AuthProxyDesc) PostParse() {
	if o.Opt.Debug || o.Opt.Log.Level == "debug" {
		log.SetLevel(log.DebugLevel)
	} else {
		lvl, err := log.ParseLevel(o.Opt.Log.Level)
		if err != nil {
			log.Errorf("error parsing loglevel: %s, using INFO", err)
			lvl = log.InfoLevel
		}
		log.SetLevel(lvl)
	}
	o.RunID = utils.MustGenerateUUID()
}

func (o *AuthProxyDesc) SaveConfig() error {
	if o.Viper == nil {
		return errors.New("viper is nil")
	}
	f, err := os.OpenFile(o.Viper.ConfigFileUsed(), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	defer func() { _ = f.Close() }()
	if err != nil {
		return err
	}
	w := bufio.NewWriter(f)
	s, _ := yaml.Marshal(o.Opt)
	_, err = w.Write(s)
	if err != nil {
		return err
	}
	_ = w.Flush()
	return nil
}

// InitCfg initConfig prepares config for the application
func InitCfg(cmd *cobra.Command, args []string) {
	printFlag, _ := cmd.Flags().GetBool("print")
	outputPath, _ := cmd.Flags().GetString("output")
	overwriteFlag, _ := cmd.Flags().GetBool("yes")

	desc := NewAuthProxyDesc()
	err := desc.Parse(cmd)
	if err != nil {
		log.Errorln(err)
		return
	}
	desc.Opt.UUID = utils.MustGenerateUUID()

	if printFlag {
		configBuffer, _ := yaml.Marshal(desc.Opt)
		fmt.Println(string(configBuffer))
	} else {
		utils.DumpOption(desc.Opt, outputPath, overwriteFlag)
	}
}

// Warm prepare the mysql connection
func Warm(cmd *cobra.Command, args []string) {
	desc := NewAuthProxyDesc()
	err := desc.Parse(cmd)
	if err != nil {
		log.Errorln(err)
		return
	}
}
