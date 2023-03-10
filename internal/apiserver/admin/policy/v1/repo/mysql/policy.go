package mysql

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	model "ingress-authproxy/internal/apiserver/admin/policy/v1/model"
	repoInterface "ingress-authproxy/internal/apiserver/admin/policy/v1/repo"
	"ingress-authproxy/internal/config"
	"ingress-authproxy/internal/utils"
)

type policyRepo struct {
	dbEngine *gorm.DB
}

func (p *policyRepo) Create(policy *model.Policy) error {
	var tmpPolicy model.Policy
	p.dbEngine.Where("name = ?", policy.Name).Find(&tmpPolicy)
	if tmpPolicy.Name != "" {
		err := errors.New("the policy with same name already exit")
		log.Errorf("%+v", err)
		return err
	}

	if err := p.dbEngine.Create(&policy).Error; err != nil {
		return err
	}

	return nil
}

func (p *policyRepo) Delete(policyName string) error {
	err := p.dbEngine.Where("name = ?", policyName).Delete(&model.Policy{}).Error
	return err
}

func (p *policyRepo) Update(policy *model.Policy) error {
	err := p.dbEngine.Save(policy).Error
	return err
}

func (p *policyRepo) Get(policyName string) (*model.Policy, error) {
	policy := &model.Policy{}
	err := p.dbEngine.Where("name= ?", policyName).First(&policy).Error
	if err != nil {
		return nil, err
	}
	return policy, nil
}

func (p *policyRepo) List() (*model.PolicyList, error) {
	ret := &model.PolicyList{}

	d := p.dbEngine.
		Order("id desc").
		Find(&ret.Items).
		Offset(-1).
		Limit(-1).
		Count(&ret.TotalCount)

	if d.Error != nil {
		return nil, d.Error
	}

	return ret, nil
}

var _ repoInterface.PolicyRepo = (*policyRepo)(nil)

// newPolicyRepo creates and returns a user storage.
func newPolicyRepo(opt *config.MySQLOpt) repoInterface.PolicyRepo {
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
	return &policyRepo{dbEngine: db}
}

func (p *policyRepo) close() error {
	dbEngine, err := p.dbEngine.DB()
	if err != nil {
		return err
	}

	return dbEngine.Close()
}
