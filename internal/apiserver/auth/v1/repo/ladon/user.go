package ladon

import (
	repoInterface "ingress-auth-proxy/internal/apiserver/admin/v1/repo"
)

type userRepo struct {
}
type UserRepo interface {
}

var _ repoInterface.UserRepo = (*userRepo)(nil)

// newUserRepo creates and returns a user storage.
func newUserRepo() repoInterface.UserRepo {
	return &userRepo{}
}
