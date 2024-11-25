package businesslogics

import (
	"architecture_template/constants/notis"
	"architecture_template/services/role/entities"
	"context"
	"errors"
)

func (tr *service) GetRoleById(id string, c context.Context) (*entities.Role, error) {
	if id == "" {
		return nil, errors.New(notis.GenericsErrorWarnMsg)
	}
	//---------------------------------------
	res, err := tr.roleRepo.GetRoleById(id, c)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, errors.New(notis.UndefinedRoleWarnMsg)
	}
	//---------------------------------------
	return res, nil
}
