package repo

import (
	model "ingress-auth-proxy/internal/apiserver/admin/user/v1/model"
)

type UserRepo interface {
	Get(username string) (*model.User, error)
	Update(user *model.User) error
}
