package externalservices

import (
	"context"
)

type IRole interface {
	GetRoleStorage(isFindId bool, c context.Context) map[string]string
}
