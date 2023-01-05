package mysql

import (
	repoInterface "ingress-auth-proxy/internal/apiserver/admin/policy/v1/repo"
	"sync"
)

type repo struct {
	policyRepo PolicyRepo
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
			policyRepo: newPolicyRepo(),
		}
	})

	return r, nil
}

func (r repo) PolicyRepo() repoInterface.PolicyRepo {
	return r.policyRepo
}

// Close closes the repo.
func (r repo) Close() error {
	return r.Close()
}
