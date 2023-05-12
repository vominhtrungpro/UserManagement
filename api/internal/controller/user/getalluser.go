package user

import (
	"context"
	"vominhtrungpro/usermanagement/internal/model"
)

func (i impl) GetAllUser(ctx context.Context) ([]model.User, error) {
	return i.repo.User().GetAllUser(ctx)
}
