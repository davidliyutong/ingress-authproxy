package repo

type Repo interface {
	PolicyRepo() PolicyRepo
	Close() error
}

var client Repo

func Client() Repo {
	return client
}

func SetClient(c Repo) {
	client = c
}
