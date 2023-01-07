package mysql

import (
	repoInterface "ingress-authproxy/internal/apiserver/admin/resource/v1/repo"
	"ingress-authproxy/internal/config"
	"sync"
)

type repo struct {
	resourceRepo repoInterface.ResourceRepo
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
			resourceRepo: newResourceRepo(opt),
		}
	})

	return r, nil
}

func (r repo) UserRepo() repoInterface.ResourceRepo {
	return r.resourceRepo
}

// Close closes the repo.
func (r repo) Close() error {
	return r.Close()
}
