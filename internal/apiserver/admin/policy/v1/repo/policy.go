package repo

import model "ingress-authproxy/internal/apiserver/admin/policy/v1/model"

type PolicyRepo interface {
	Create(policy *model.Policy) error
	Delete(policyName string) error
	Update(policy *model.Policy) error
	Get(policyName string) (*model.Policy, error)
	List() (*model.PolicyList, error)
}
