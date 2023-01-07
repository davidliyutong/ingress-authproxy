package mysql

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	model "ingress-authproxy/internal/apiserver/admin/user/v1/model"
	repoInterface "ingress-authproxy/internal/apiserver/admin/user/v1/repo"
	"ingress-authproxy/internal/config"
	"ingress-authproxy/internal/utils"
	"regexp"
)

type userRepo struct {
	dbEngine *gorm.DB
}

func (u *userRepo) Create(user *model.User) error {
	tmpUser := model.User{}
	u.dbEngine.Where("name = ?", user.Name).Find(&tmpUser)
	if tmpUser.Name != "" {
		err := errors.New("the created user already exit")

		log.Errorf("%+v", err)
		return err
	}

	err := u.dbEngine.Create(&user).Error
	if err != nil {
		if match, _ := regexp.MatchString("Duplicate entry", err.Error()); match {
			return errors.New("duplicate entry")
		}

		return err
	}

	return nil
}

func (u *userRepo) Delete(username string) error {
	if err := u.dbEngine.Where("name = ?", username).Delete(&model.User{}).Error; err != nil {
		return err
	}

	return nil
}

func (u *userRepo) List() (*model.UserList, error) {
	ret := &model.UserList{}

	d := u.dbEngine.
		Order("id desc").
		Find(&ret.Items).
		Offset(-1).
		Limit(-1).
		Count(&ret.TotalCount)

	return ret, d.Error
}

func (u *userRepo) Update(user *model.User) error {
	tmpUser := model.User{}
	u.dbEngine.Where("name = ?", user.Name).Find(&tmpUser)
	if tmpUser.Name == "" {
		err := errors.New("the update user not found")
		log.Errorf("%s\n", err)
		return err
	}

	if err := u.dbEngine.Save(user).Error; err != nil {
		return err
	}

	return nil
}

func (u *userRepo) Get(username string) (*model.User, error) {
	user := &model.User{}
	err := u.dbEngine.Where("name = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(fmt.Sprintf("the user: %v not found", username))
		}
		return nil, errors.New(err.Error())
	}

	return user, nil
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

func (p *userRepo) close() error {
	dbEngine, err := p.dbEngine.DB()
	if err != nil {
		return err
	}

	return dbEngine.Close()
}
