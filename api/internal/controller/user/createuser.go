package user

import (
	"context"
	"vominhtrungpro/usermanagement/internal/model"
)

type CreateInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Age      int64  `json:"age"`
}

func (i impl) CreateUser(ctx context.Context, input CreateInput) (model.User, error) {
	return i.repo.User().CreateUser(ctx, model.User{
		Username: input.Username,
		Password: input.Password,
		Email:    input.Email,
		Age:      input.Age,
	})
}
