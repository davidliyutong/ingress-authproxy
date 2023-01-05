package mysql

import (
	repo2 "ingress-auth-proxy/internal/apiserver/admin/v1/repo"
	"sync"
)

type repo struct {
	userRepo   UserRepo
	secretRepo SecretRepo
	policyRepo PolicyRepo
}

//var _ repo3.BlobRepo = (*repo)(nil)

var (
	r    repo
	once sync.Once
)

// Repo creates and returns the store client instance.
func Repo() (repo2.Repo, error) {
	once.Do(func() {
		r = repo{
			userRepo:   newUserRepo(),
			secretRepo: newSecretRepo(),
			policyRepo: newPolicyRepo(),
		}
	})

	return r, nil
}

func (r repo) UserRepo() repo2.UserRepo {
	return r.userRepo
}

func (r repo) SecretRepo() repo2.SecretRepo {
	return r.secretRepo
}

func (r repo) PolicyRepo() repo2.PolicyRepo {
	return r.policyRepo
}

// Close closes the repo.
func (r repo) Close() error {
	return r.Close()
}
