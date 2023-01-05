package mysql

import (
	"fmt"
	"github.com/rebirthmonkey/go/pkg/errcode"
	"github.com/rebirthmonkey/go/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	model "ingress-auth-proxy/internal/apiserver/admin/user/v1/model"
	repoInterface "ingress-auth-proxy/internal/apiserver/admin/user/v1/repo"
	"ingress-auth-proxy/internal/config"
	"ingress-auth-proxy/internal/utils"
)

type userRepo struct {
	dbEngine *gorm.DB
}

func (u *userRepo) Update(user *model.User) error {
	tmpUser := model.User{}
	u.dbEngine.Where("name = ?", user.Name).Find(&tmpUser)
	if tmpUser.Name == "" {
		err := errors.WithCode(errcode.ErrRecordNotFound, "the update user not found")
		log.Errorf("%s\n", err)
		return err
	}

	if err := u.dbEngine.Save(user).Error; err != nil {
		return err
	}

	return nil
}

func (u userRepo) Get(username string) (*model.User, error) {
	user := &model.User{}
	err := u.dbEngine.Where("name = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.WithCode(errcode.ErrRecordNotFound, "the get user not found.")
		}
		return nil, errors.WithCode(errcode.ErrDatabase, err.Error())
	}

	return user, nil
}

type UserRepo interface {
	Get(username string) (*model.User, error)
	Update(user *model.User) error
}

var _ repoInterface.UserRepo = (*userRepo)(nil)

// newUserRepo creates and returns a user storage.
func newUserRepo(opt *config.MySQLOpt) repoInterface.UserRepo {
	dsn := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s`,
		opt.Username,
		opt.Password,
		opt.Hostname,
		opt.Database,
		true,
		utils.GetMySQLTZFromEnv())

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Mysql connection fails %+v\n", err)
		return nil
	}
	return &userRepo{dbEngine: db}
}
