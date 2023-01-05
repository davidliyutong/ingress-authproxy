package mysql

import (
	repoInterface "ingress-auth-proxy/internal/apiserver/admin/user/v1/repo"
	"ingress-auth-proxy/internal/config"
	"sync"
)

type repo struct {
	userRepo UserRepo
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
			userRepo: newUserRepo(opt),
		}
	})

	return r, nil
}

func (r repo) UserRepo() repoInterface.UserRepo {
	return r.userRepo
}

// Close closes the repo.
func (r repo) Close() error {
	return r.Close()
}
