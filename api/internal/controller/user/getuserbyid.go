package user

import (
	"context"
	"vominhtrungpro/usermanagement/internal/model"
)

func (i impl) GetUserByID(ctx context.Context, id int64) (model.User, error) {
	return i.repo.User().GetUserByID(ctx, id)
}
