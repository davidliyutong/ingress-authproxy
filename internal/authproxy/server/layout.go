package server

type LayoutAdmin struct {
	Self     string
	Users    string
	Policies string
	Secrets  string
	Info     string
	Server   LayoutServer
}

type LayoutJWT struct {
	Self    string
	Login   string
	Refresh string
}

type LayoutIngressAuth struct {
	Self string
}

type LayoutServer struct {
	Self     string
	Option   string
	Shutdown string
	Sync     string
}

//type LayoutUser struct {
//	Self string
//}

type LayoutV1 struct {
	Self        string
	Admin       LayoutAdmin
	User        string
	IngressAuth LayoutIngressAuth
	JWT         LayoutJWT
	Server      LayoutServer
	Healthz     string
	Version     string
}

type Layout struct {
	Ping    string
	V1      LayoutV1
	Healthz string
	Version string
}

var AuthProxyLayout = Layout{
	Ping: "/ping",

	V1: LayoutV1{
		Self: "/v1",
		Admin: LayoutAdmin{
			Self:     "/v1/admin",
			Users:    "users",
			Policies: "policies",
			Secrets:  "secrets",
			Info:     "info",
			Server: LayoutServer{
				Self:     "/v1/admin/server",
				Option:   "option",
				Shutdown: "shutdown",
				Sync:     "sync",
			},
		},

		User: "user",

		IngressAuth: LayoutIngressAuth{
			Self: "/v1/ingress-auth",
		},
		JWT: LayoutJWT{
			Self:    "/v1/jwt",
			Login:   "login",
			Refresh: "refresh",
		},
		Healthz: "healthz",
		Version: "version",
	},
	Healthz: "healthz",
	Version: "version",
}
