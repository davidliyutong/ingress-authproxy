package server

type LayoutAdmin struct {
	Self     string
	Users    string
	Policies string
	Secrets  string
	Info     string
}

type LayoutJWT struct {
	Self    string
	Login   string
	Refresh string
}

type LayoutIngressAuth struct {
	Self string
}

type LayoutV1 struct {
	Self        string
	Admin       LayoutAdmin
	IngressAuth LayoutIngressAuth
	JWT         LayoutJWT
}

type Layout struct {
	Ping string
	V1   LayoutV1
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
		},

		IngressAuth: LayoutIngressAuth{
			Self: "/v1/ingress-auth",
		},
		JWT: LayoutJWT{
			Self:    "/v1/jwt",
			Login:   "login",
			Refresh: "refresh",
		},
	},
}
