package ladon

import (
	repoInterface "ingress-auth-proxy/internal/apiserver/auth/v1/repo"
	"sync"
)

type repo struct {
	authRepo repoInterface.AuthRepo
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
			authRepo: newUserRepo(),
		}
	})

	return r, nil
}

func (r repo) AuthRepo() repoInterface.AuthRepo {
	return r.authRepo
}

// Close closes the repo.
func (r repo) Close() error {
	return r.Close()
}
