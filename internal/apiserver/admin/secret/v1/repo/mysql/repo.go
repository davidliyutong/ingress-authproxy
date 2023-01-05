package mysql

import (
	repoInterface "ingress-auth-proxy/internal/apiserver/admin/secret/v1/repo"
	"sync"
)

type repo struct {
	secretRepo SecretRepo
}

//var _ repo3.BlobRepo = (*repo)(nil)

var (
	r    repo
	once sync.Once
)

// Repo creates and returns the store client instance.
func Repo() (repoInterface.Repo, error) {
	once.Do(func() {
		r = repo{
			secretRepo: newSecretRepo(),
		}
	})

	return r, nil
}

func (r repo) SecretRepo() repoInterface.SecretRepo {
	return r.secretRepo
}

// Close closes the repo.
func (r repo) Close() error {
	return r.Close()
}
