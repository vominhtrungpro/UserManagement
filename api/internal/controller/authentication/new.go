package authentication

import (
	"context"
	"vominhtrungpro/usermanagement/internal/repository"
)

type Controller interface {
	Login(context.Context, LoginInput) (TokenOutput, error)
}

func New(repo repository.Registry) Controller {
	return impl{
		repo: repo,
	}
}

type impl struct {
	repo repository.Registry
}
