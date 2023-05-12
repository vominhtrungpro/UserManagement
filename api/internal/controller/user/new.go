package user

import (
	"context"
	"vominhtrungpro/usermanagement/internal/model"
	"vominhtrungpro/usermanagement/internal/repository"
)

type Controller interface {
	GetAllUser(context.Context) ([]model.User, error)
	CreateUser(context.Context, CreateInput) (model.User, error)
	UpdateUser(context.Context, UpdateInput) (model.User, error)
	DeleteUser(context.Context, int64) error
	GetUserByID(context.Context, int64) (model.User, error)
}

func New(repo repository.Registry) Controller {
	return impl{
		repo: repo,
	}
}

type impl struct {
	repo repository.Registry
}
