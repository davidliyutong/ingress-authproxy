package ladon

import (
	repoInterface "ingress-auth-proxy/internal/apiserver/auth/v1/repo"
	"sync"
)

type repo struct {
	authzRepo repoInterface.AuthzRepo
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
			authzRepo: newUserRepo(),
		}
	})

	return r, nil
}

func (r repo) AuthzRepo() repoInterface.AuthzRepo {
	return r.authzRepo
}

// Close closes the repo.
func (r repo) Close() error {
	return r.Close()
}
