package user

import (
	"context"
	"database/sql"
	"vominhtrungpro/usermanagement/internal/model"
)

type Repository interface {
	// specification
	GetAllUser(context.Context) ([]model.User, error)
}

func New(db *sql.DB) Repository {
	return impl{
		db: db,
	}
}

type impl struct {
	db *sql.DB
}
