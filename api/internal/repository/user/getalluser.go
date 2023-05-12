package user

import (
	"context"
	"vominhtrungpro/usermanagement/internal/model"
	dbmodel "vominhtrungpro/usermanagement/internal/repository/dbmodel"

	pkgerrors "github.com/pkg/errors"
)

func (i impl) GetAllUser(ctx context.Context) ([]model.User, error) {
	usersdb, err := dbmodel.Users().All(ctx, i.db)
	if err != nil {
		return []model.User{}, pkgerrors.WithStack(err)
	}
	var users []model.User
	for _, item := range usersdb {
		var user model.User
		user.ID = item.ID
		user.Username = item.Username
		user.Password = item.Password
		user.Email = item.Email
		user.Age = int64(item.Age)
		users = append(users, user)
	}
	return users, nil
}
