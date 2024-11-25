package businesslogics

import (
	"architecture_template/services/user/entities"
	"context"
)

func (tr *service) GetAllUsers(c context.Context) (*[]entities.User, error) {
	return tr.repo.GetAllUsers(c)
}
