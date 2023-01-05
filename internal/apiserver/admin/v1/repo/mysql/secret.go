package mysql

import (
	repoInterface "ingress-auth-proxy/internal/apiserver/admin/v1/repo"
)

type secretRepo struct {
}
type SecretRepo interface {
}

var _ repoInterface.SecretRepo = (*secretRepo)(nil)

// newSecretRepo creates and returns a user storage.
func newSecretRepo() repoInterface.SecretRepo {
	return &secretRepo{}
}
