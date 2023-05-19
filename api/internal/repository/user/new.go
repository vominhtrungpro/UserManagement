package user

import (
	"context"
	"database/sql"
	"vominhtrungpro/usermanagement/internal/model"
)

type Repository interface {
	// specification
	GetAllUser(context.Context) ([]model.User, error)
	CreateUser(context.Context, model.User) (model.User, error)
	UpdateUser(context.Context, model.User) (model.User, error)
	DeleteUser(context.Context, int64) error
	GetUserByID(context.Context, int64) (model.User, error)
	GetUserByUsername(context.Context, string) (model.User, error)
	UpdateRefreshToken(context.Context, model.User, string) error
}

func New(db *sql.DB) Repository {
	return impl{
		db: db,
	}
}

type impl struct {
	db *sql.DB
}
