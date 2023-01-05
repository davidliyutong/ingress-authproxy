package config

import (
	"bufio"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/yaml.v3"
	"ingress-auth-proxy/internal/utils"
	"os"
	"path"
	"strings"
	"sync"
	"time"
)

var userHomeDir, _ = os.UserHomeDir()

const DefaultDomain = "localhost"
const DefaultPort = 50032
const DefaultInterface = "0.0.0.0"
const DefaultAppName = "authproxy"
const DefaultConfigName = "nameserver"
const DefaultTimeoutSecond = 2 * 3600
const DefaultInitUsername = "admin"
const DefaultInitPassword = "admin"

var DefaultConfig = path.Join(userHomeDir, ".config/"+DefaultAppName+"/"+DefaultConfigName+".yaml")

const DefaultConfigSearchPath0 = "/etc/authproxy"
const DefaultConfigSearchPath1 = "./"
const DefaultConfigSearchPath2 = "/config"

var (
	GlobalServerDesc *AuthProxyDesc
	once             sync.Once
)

// SetGlobalDesc set the global desc for once
func SetGlobalDesc(desc *AuthProxyDesc) {
	once.Do(func() {
		GlobalServerDesc = desc
	})
}

type MySQLOpt struct {
	Hostname    string `yaml:"hostname"`
	Port        int    `yaml:"port"`
	Database    string `yaml:"database"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	Initialized bool   `yaml:"initialized"`
}

type NetworkOpt struct {
	Port      int    `yaml:"port"`
	Interface string `yaml:"interface"`
	Domain    string `yaml:"domain"`
}

type LogOpt struct {
	Level string `yaml:"port"`
	Path  string `yaml:"path"`
}

type JWTOpt struct {
	Secret  string `yaml:"-"`
	Timeout string `yaml:"timeout"`
}

// InitOpt init the admin user
type InitOpt struct {
	Username string `yaml:"-"`
	Password string `yaml:"-"`
}
type AuthProxyOpt struct {
	UUID    string     `yaml:"uuid"`
	Network NetworkOpt `yaml:"network"`
	Debug   bool       `yaml:"debug"`
	Log     LogOpt     `yaml:"log"`
	MySQL   MySQLOpt   `yaml:"mysql"`
	JWT     JWTOpt     `yaml:"jwt"`
	Init    InitOpt    `yaml:"-"`
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
			Domain:    DefaultDomain,
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
		Init: InitOpt{
			Username: "",
			Password: "",
		},
	}
}

func (o *AuthProxyDesc) Parse(cmd *cobra.Command) error {
	vipCfg := viper.New()
	vipCfg.SetDefault("network.domain", DefaultDomain)
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
	vipCfg.SetDefault("jwt.expired", DefaultTimeoutSecond*time.Second)
	vipCfg.SetDefault("init.username", DefaultInitPassword)
	vipCfg.SetDefault("init.password", DefaultInitPassword)

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
	if o.Opt.UUID == "" {
		o.Opt.UUID = utils.MustGenerateUUID()
	}
	if o.Opt.JWT.Secret == "" {
		_, s := utils.MustGenerateAuthKeys()
		o.Opt.JWT.Secret = s
	}
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

const warmSQLString = `
DROP TABLE IF EXISTS user;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE user (
id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
instanceID varchar(32) DEFAULT NULL,
name varchar(45) NOT NULL,
status int(1) DEFAULT 1 COMMENT '1:可用，0:不可用',
nickname varchar(30) NOT NULL,
password varchar(255) NOT NULL,
email varchar(256) NOT NULL,
phone varchar(20) DEFAULT NULL,
isAdmin tinyint(1) unsigned NOT NULL DEFAULT 0 COMMENT '1: administrator\\\\n0: non-administrator',
extendShadow longtext DEFAULT NULL,
loginedAt timestamp NULL DEFAULT NULL COMMENT 'last login time',
createdAt timestamp NOT NULL DEFAULT current_timestamp(),
updatedAt timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
PRIMARY KEY (id),
UNIQUE KEY idx_name (name),
UNIQUE KEY instanceID_UNIQUE (instanceID)
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table secret
--

DROP TABLE IF EXISTS secret;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE secret (
id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
instanceID varchar(32) DEFAULT NULL,
name varchar(45) NOT NULL,
username varchar(255) NOT NULL,
secretID varchar(36) NOT NULL,
secretKey varchar(255) NOT NULL,
expires int(64) unsigned NOT NULL DEFAULT 1534308590,
description varchar(255) NOT NULL,
extendShadow longtext DEFAULT NULL,
createdAt timestamp NOT NULL DEFAULT current_timestamp(),
updatedAt timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
PRIMARY KEY (id),
UNIQUE KEY instanceID_UNIQUE (instanceID)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table policy
--

DROP TABLE IF EXISTS policy;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE policy (
id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
instanceID varchar(32) DEFAULT NULL,
name varchar(45) NOT NULL,
username varchar(255) NOT NULL,
policyShadow longtext DEFAULT NULL,
extendShadow longtext DEFAULT NULL,
createdAt timestamp NOT NULL DEFAULT current_timestamp(),
updatedAt timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
PRIMARY KEY (id),
UNIQUE KEY instanceID_UNIQUE (instanceID)
) ENGINE=InnoDB AUTO_INCREMENT=47 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;`

// Warm prepare the mysql connection
func Warm(db *sql.DB) error {
	_, err := db.Query(warmSQLString)
	return err
}

// Verify check the mysql database
func Verify(db *sql.DB) bool {
	_, check0 := db.Query("select * from " + "user" + ";")
	log.Debugf("table user exist: %t", check0 == nil)
	_, check1 := db.Query("select * from " + "policy" + ";")
	log.Debugf("table policy exist: %t", check0 == nil)
	_, check2 := db.Query("select * from " + "secret" + ";")
	log.Debugf("table secret exist: %t", check0 == nil)

	if check0 == nil && check1 == nil && check2 == nil {
		return true
	} else {
		return false
	}
}

func CreateDefaultAdminUser(db *sql.DB, username string, password string) error {
	// check if the user table is empty
	var count int
	err := db.QueryRow("select count(*) from user;").Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		log.Infoln("user table is not empty, skip the default admin user creation")
		return nil
	} else {
		log.Infoln("user table is empty, create the default admin user")
		encrypted, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		sqlString := fmt.Sprintf("insert into `user` values (%d, '%s', '%s', %d, '%s', '%s', '%s', '%s', %d, '%s', now(), '%s', '%s');",
			0, "user", username, 1, "admin", encrypted, "admin@example.com", "+0000000000000", 1, "{}", utils.GetMySQLTime(time.Now()), utils.GetMySQLTime(time.Now()))
		log.Debugf("sqlString: %s", sqlString)
		_, err := db.Exec(sqlString)
		return err
	}
}
