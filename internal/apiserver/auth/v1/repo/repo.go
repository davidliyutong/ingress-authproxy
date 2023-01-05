package repo

type Repo interface {
	AuthRepo() AuthRepo
	Close() error
}

var client Repo

func Client() Repo {
	return client
}

func SetClient(c Repo) {
	client = c
}