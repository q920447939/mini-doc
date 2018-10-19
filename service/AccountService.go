package service

import (
	"wahaha/module/rbac"
	"wahaha/base"
)

type  AccountService interface {
	AddMember(m *rbac.Member) (e *	base.BaseReturnJson)
	Login(m *rbac.Member) (e *	base.BaseReturnJson)
}
