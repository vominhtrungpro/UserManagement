package user

import (
	"context"
	"vominhtrungpro/usermanagement/internal/model"
	"vominhtrungpro/usermanagement/internal/repository"
)

type Controller interface {
	GetAllUser(context.Context) ([]model.User, error)
}

func New(repo repository.Registry) Controller {
	return impl{
		repo: repo,
	}
}

type impl struct {
	repo repository.Registry
}
