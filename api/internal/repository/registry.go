package repository

import (
	"database/sql"
	"vominhtrungpro/usermanagement/internal/repository/user"
)

type Registry interface {
	// Inventory returns the inventory repo
	User() user.Repository
}

func New(db *sql.DB) Registry {
	return impl{
		user: user.New(db),
	}
}

type impl struct {
	user user.Repository
}

func (i impl) User() user.Repository {
	return i.user
}
