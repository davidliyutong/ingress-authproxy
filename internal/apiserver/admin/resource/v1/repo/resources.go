package repo

import (
	model "ingress-authproxy/internal/apiserver/admin/resource/v1/model"
)

type ResourceRepo interface {
	Create(user *model.Resource) error
	Delete(username string) error
	Update(user *model.Resource) error
	Get(username string) (*model.Resource, error)
	List() (*model.ResourceList, error)
}
