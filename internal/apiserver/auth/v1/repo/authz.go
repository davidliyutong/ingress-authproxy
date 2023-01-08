package repo

import (
	"github.com/ory/ladon"
	policyRepo "ingress-authproxy/internal/apiserver/admin/policy/v1/repo"
	userRepo "ingress-authproxy/internal/apiserver/admin/user/v1/repo"
)

type AuthzRepo interface {
	List() ([]*ladon.DefaultPolicy, error)
	Start()
	Stop()
	Trigger()
	Validate(request *ladon.Request) error
	UserRepo() userRepo.UserRepo
	PolicyRepo() policyRepo.PolicyRepo
}
