package mysql

import (
	repoInterface "ingress-authproxy/internal/apiserver/admin/policy/v1/repo"
	"ingress-authproxy/internal/config"
	"sync"
)

type repo struct {
	policyRepo repoInterface.PolicyRepo
}

//var _ repo3.BlobRepo = (*repo)(nil)

var (
	r    repo
	once sync.Once
)

// Repo creates and returns the store client instance.
func Repo(opt *config.MySQLOpt) (repoInterface.Repo, error) {
	once.Do(func() {
		r = repo{
			policyRepo: newPolicyRepo(opt),
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
