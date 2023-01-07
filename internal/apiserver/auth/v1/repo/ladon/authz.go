package ladon

import (
	repoInterface "ingress-authproxy/internal/apiserver/auth/v1/repo"
)

type authzRepo struct {
}
type AuthzRepo interface {
}

var _ repoInterface.AuthzRepo = (*authzRepo)(nil)

// newUserRepo creates and returns a user storage.
func newUserRepo() repoInterface.AuthzRepo {
	return &authzRepo{}
}
