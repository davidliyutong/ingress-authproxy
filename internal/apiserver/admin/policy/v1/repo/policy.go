package repo

import model "ingress-authproxy/internal/apiserver/admin/policy/v1/model"

type PolicyRepo interface {
	Create(policy *model.Policy) error
	Delete(username string, policyName string) error
	DeleteByUser(username string) error
	Update(policy *model.Policy) error
	Get(username string, policyName string) (*model.Policy, error)
	List(username string) (*model.PolicyList, error)
}
