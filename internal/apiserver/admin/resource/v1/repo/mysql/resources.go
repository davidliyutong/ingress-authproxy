package mysql

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	model "ingress-authproxy/internal/apiserver/admin/resource/v1/model"
	repoInterface "ingress-authproxy/internal/apiserver/admin/resource/v1/repo"
	"ingress-authproxy/internal/config"
	"ingress-authproxy/internal/utils"
	"regexp"
)

type resourceRepo struct {
	dbEngine *gorm.DB
}

func (u *resourceRepo) Create(user *model.Resource) error {
	tmpUser := model.Resource{}
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

func (u *resourceRepo) Delete(username string) error {
	if err := u.dbEngine.Where("name = ?", username).Delete(&model.Resource{}).Error; err != nil {
		return err
	}

	return nil
}

func (u *resourceRepo) List() (*model.ResourceList, error) {
	ret := &model.ResourceList{}

	d := u.dbEngine.
		Order("id desc").
		Find(&ret.Items).
		Offset(-1).
		Limit(-1).
		Count(&ret.TotalCount)

	return ret, d.Error
}

func (u *resourceRepo) Update(user *model.Resource) error {
	tmpUser := model.Resource{}
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

func (u *resourceRepo) Get(username string) (*model.Resource, error) {
	user := &model.Resource{}
	err := u.dbEngine.Where("name = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(fmt.Sprintf("the user: %v not found", username))
		}
		return nil, errors.New(err.Error())
	}

	return user, nil
}

var _ repoInterface.ResourceRepo = (*resourceRepo)(nil)

// newResourceRepo creates and returns a user storage.
func newResourceRepo(opt *config.MySQLOpt) repoInterface.ResourceRepo {
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
	return &resourceRepo{dbEngine: db}
}
