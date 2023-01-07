package v1

import (
	"golang.org/x/crypto/bcrypt"
	model "ingress-authproxy/internal/apiserver/admin/resource/v1/model"
	"ingress-authproxy/internal/apiserver/admin/resource/v1/repo"
	"ingress-authproxy/pkg/metamodel"
	"time"
)

type ResourceService interface {
	Create(user *model.Resource) error
	Delete(username string) error
	Update(user *model.Resource) error
	Get(username string) (*model.Resource, error)
	List() (*model.ResourceList, error)
}

type resourceService struct {
	repo repo.Repo
}

// Create creates a new user account.
func (u *resourceService) Create(user *model.Resource) error {
	hashedBytes, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedBytes)
	user.Status = 1
	user.LoginedAt = time.Now()

	return u.repo.UserRepo().Create(user)
}

// Delete deletes the user by the user identifier.
func (u *resourceService) Delete(username string) error {
	return u.repo.UserRepo().Delete(username)
}

// Update updates a user account information.
func (u *resourceService) Update(user *model.Resource) error {
	updateUser, err := u.Get(user.Name)
	if err != nil {
		return err
	}

	updateUser.Status = user.Status
	updateUser.Nickname = user.Nickname
	hashedBytes, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	updateUser.Password = string(hashedBytes)
	updateUser.Email = user.Email
	updateUser.Phone = user.Phone
	updateUser.Extend = user.Extend
	updateUser.IsAdmin = user.IsAdmin

	return u.repo.UserRepo().Update(updateUser)
}

// Get returns a user's info by the user identifier.
func (u *resourceService) Get(username string) (*model.Resource, error) {
	return u.repo.UserRepo().Get(username)
}

// List returns all the related users.
func (u *resourceService) List() (*model.ResourceList, error) {
	users, err := u.repo.UserRepo().List()
	if err != nil {
		return nil, err
	}

	infos := make([]*model.Resource, 0)
	for _, user := range users.Items {
		infos = append(infos, &model.Resource{
			ObjectMeta: metamodel.ObjectMeta{
				ID:        user.ID,
				Name:      user.Name,
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.UpdatedAt,
			},
			Nickname: user.Nickname,
			Email:    user.Email,
			Phone:    user.Phone,
		})
	}

	return &model.ResourceList{ListMeta: users.ListMeta, Items: infos}, nil
}

func newAdminService(repo repo.Repo) ResourceService {
	return &resourceService{repo: repo}
}

func (s *service) NewResourceService() ResourceService {
	return newAdminService(s.repo)
}

var _ ResourceService = (*resourceService)(nil)
