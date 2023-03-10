package v1

import (
	"golang.org/x/crypto/bcrypt"
	model "ingress-authproxy/internal/apiserver/admin/user/v1/model"
	"ingress-authproxy/internal/apiserver/admin/user/v1/repo"
	"ingress-authproxy/pkg/metamodel"
	"time"
)

type UserService interface {
	Create(user *model.User) error
	Delete(username string) error
	Update(user *model.User) error
	Get(username string) (*model.User, error)
	List() (*model.UserList, error)
}

type userService struct {
	repo repo.Repo
}

// Create creates a new user account.
func (u *userService) Create(user *model.User) error {
	hashedBytes, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedBytes)
	user.Status = 1
	user.LoginedAt = time.Now()
	if user.Name == "" {
		user.Name = user.Nickname
	}

	return u.repo.UserRepo().Create(user)
}

// Delete deletes the user by the user identifier.
func (u *userService) Delete(username string) error {
	return u.repo.UserRepo().Delete(username)
}

// Update updates a user account information.
func (u *userService) Update(user *model.User) error {
	updateUser, err := u.Get(user.Name)
	if err != nil {
		return err
	}

	updateUser.Status = user.Status
	updateUser.Nickname = user.Nickname
	if user.Password != "" {
		hashedBytes, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		updateUser.Password = string(hashedBytes)
	}
	updateUser.Email = user.Email
	updateUser.Phone = user.Phone
	updateUser.Extend = user.Extend
	updateUser.IsAdmin = user.IsAdmin

	return u.repo.UserRepo().Update(updateUser)
}

// Get returns a user's info by the user identifier.
func (u *userService) Get(username string) (*model.User, error) {
	return u.repo.UserRepo().Get(username)
}

// List returns all the related users.
func (u *userService) List() (*model.UserList, error) {
	users, err := u.repo.UserRepo().List()
	if err != nil {
		return nil, err
	}

	infos := make([]*model.User, 0)
	for _, user := range users.Items {
		infos = append(infos, &model.User{
			ObjectMeta: metamodel.ObjectMeta{
				ID:        user.ID,
				Name:      user.Name,
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.UpdatedAt,
			},
			Nickname:  user.Nickname,
			Email:     user.Email,
			Phone:     user.Phone,
			IsAdmin:   user.IsAdmin,
			Status:    user.Status,
			LoginedAt: user.LoginedAt,
		})
	}

	return &model.UserList{ListMeta: users.ListMeta, Items: infos}, nil
}

func newAdminService(repo repo.Repo) UserService {
	return &userService{repo: repo}
}

func (s *service) NewUserService() UserService {
	return newAdminService(s.repo)
}

var _ UserService = (*userService)(nil)
