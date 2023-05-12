package user

import (
	"context"
	"vominhtrungpro/usermanagement/internal/model"
)

type UpdateInput struct {
	Id       int64
	Username string
	Password string
	Email    string
	Age      int64
}

func (i impl) UpdateUser(ctx context.Context, input UpdateInput) (model.User, error) {
	return i.repo.User().UpdateUser(ctx, model.User{
		ID:       input.Id,
		Username: input.Username,
		Password: input.Password,
		Email:    input.Email,
		Age:      input.Age,
	})
}
