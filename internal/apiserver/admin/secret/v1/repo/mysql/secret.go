package mysql

import (
	repoInterface "ingress-authproxy/internal/apiserver/admin/secret/v1/repo"
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
