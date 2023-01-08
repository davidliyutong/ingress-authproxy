package ladon

import (
	policyRepo "ingress-authproxy/internal/apiserver/admin/policy/v1/repo"
	userRepo "ingress-authproxy/internal/apiserver/admin/user/v1/repo"
	repoInterface "ingress-authproxy/internal/apiserver/auth/v1/repo"
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
func Repo(userRepoClient userRepo.UserRepo, policyRepoClient policyRepo.PolicyRepo) (repoInterface.Repo, error) {
	once.Do(func() {
		r = repo{
			authzRepo: newAuthzRepo(userRepoClient, policyRepoClient),
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
