package repo

import (
	model "ingress-auth-proxy/internal/apiserver/admin/user/v1/model"
)

type UserRepo interface {
	Create(user *model.User) error
	Delete(username string) error
	Update(user *model.User) error
	Get(username string) (*model.User, error)
	List() (*model.UserList, error)
}
