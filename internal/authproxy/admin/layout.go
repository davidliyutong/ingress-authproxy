package admin

type LayoutAdmin struct {
	Self     string
	Users    string
	Policies string
	Secrets  string
	Info     string
}

type LayoutAuth struct {
	Self string
}

type LayoutV1 struct {
	Self  string
	Admin LayoutAdmin
	Auth  LayoutAuth
}

type Layout struct {
	V1 LayoutV1
}

var AuthProxyLayout = Layout{
	V1: LayoutV1{
		Self: "/v1",
		Admin: LayoutAdmin{
			Self:     "/v1/admin",
			Users:    "users",
			Policies: "policies",
			Secrets:  "secrets",
			Info:     "info",
		},
		Auth: LayoutAuth{
			Self: "/v1/auth",
		},
	},
}
