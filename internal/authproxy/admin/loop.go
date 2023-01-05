package admin

import log "github.com/sirupsen/logrus"

func RunFunc(a AuthProxyAdmin) {
	aVal := a.(*authProxyAdmin)
	log.Infoln("uuid:", aVal.desc.Opt.UUID)
	log.Infoln("port:", aVal.desc.Opt.Network.Port)
	log.Infoln("interface:", aVal.desc.Opt.Network.Interface)
	log.Infoln("mysql.hostname:", aVal.desc.Opt.MySQL.Hostname)
	log.Infoln("mysql.port:", aVal.desc.Opt.MySQL.Port)
	log.Infoln("mysql.database:", aVal.desc.Opt.MySQL.Database)
}
