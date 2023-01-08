package v1

import (
	"golang.org/x/crypto/bcrypt"
	model "ingress-authproxy/internal/apiserver/admin/user/v1/model"
	"ingress-authproxy/internal/apiserver/admin/user/v1/repo"
)

type UserService interface {
	Update(user *model.User) error
	Get(username string) (*model.User, error)
}

type userService struct {
	repo repo.UserRepo
}

// Update updates a user account information.
func (u *userService) Update(user *model.User) error {
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

	return u.repo.Update(updateUser)
}

// Get returns a user's info by the user identifier.
func (u *userService) Get(username string) (*model.User, error) {
	return u.repo.Get(username)
}

func newAdminService(repo repo.UserRepo) UserService {
	return &userService{repo: repo}
}

func (s *service) NewUserService() UserService {
	return newAdminService(s.repo)
}

var _ UserService = (*userService)(nil)
