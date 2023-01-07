package main

import (
	"database/sql"
	"fmt"
	log "github.com/sirupsen/logrus"
	"ingress-authproxy/internal/config"
	"ingress-authproxy/internal/utils"
	"os"
)

func main() {
	log.SetLevel(log.DebugLevel)
	opt := config.MySQLOpt{
		Hostname:    "localhost",
		Port:        3306,
		Database:    "authproxy",
		Username:    "authproxy",
		Password:    "authproxy",
		Initialized: false,
	}
	dsn := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&multiStatements=true&loc=%s`,
		opt.Username,
		opt.Password,
		opt.Hostname,
		opt.Database,
		utils.GetMySQLTZFromEnv())

	db, err := sql.Open("mysql",
		dsn)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	if !opt.Initialized {
		// do something
		log.Println(config.Verify(db))
		err := config.Warm(db)
		if err != nil {
			log.Error(err)
			os.Exit(1)
		}
		opt.Initialized = config.Verify(db)
		log.Println(opt.Initialized)

	}
}
