package repo

type Repo interface {
	SecretRepo() SecretRepo
	Close() error
}

var client Repo

func Client() Repo {
	return client
}

func SetClient(c Repo) {
	client = c
}
