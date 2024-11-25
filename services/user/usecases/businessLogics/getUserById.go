package businesslogics

import (
	"architecture_template/constants/notis"
	"architecture_template/services/user/entities"
	"context"
	"errors"
	"strings"
)

func (tr *service) GetUserById(id string, c context.Context) (*entities.User, error) {
	if id = strings.TrimSpace(id); id != "" {
		return tr.repo.GetUserById(id, c)
	}

	return nil, errors.New(notis.GenericsErrorWarnMsg)
}
