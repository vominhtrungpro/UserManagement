package repository

import (
	"database/sql"
	"vominhtrungpro/usermanagement/internal/repository/authentication"
	"vominhtrungpro/usermanagement/internal/repository/user"
)

type Registry interface {
	User() user.Repository
	Authentication() authentication.Repository
}

func New(db *sql.DB) Registry {
	return impl{
		user:           user.New(db),
		authentication: authentication.New(db),
	}
}

type impl struct {
	user           user.Repository
	authentication authentication.Repository
}

func (i impl) User() user.Repository {
	return i.user
}

func (i impl) Authentication() authentication.Repository {
	return i.authentication
}
