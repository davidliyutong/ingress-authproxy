package mysql

import (
	repoInterface "ingress-auth-proxy/internal/apiserver/admin/v1/repo"
)

type policyRepo struct {
}
type PolicyRepo interface {
}

var _ repoInterface.PolicyRepo = (*policyRepo)(nil)

// newPolicyRepo creates and returns a user storage.
func newPolicyRepo() repoInterface.PolicyRepo {
	return &policyRepo{}
}
