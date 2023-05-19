package authentication

import (
	"context"
	"errors"
	"vominhtrungpro/usermanagement/internal/model"
	dbmodel "vominhtrungpro/usermanagement/internal/repository/dbmodel"
)

func (i impl) Login(ctx context.Context, username string, password string) (model.User, error) {
	userdb, err := dbmodel.Users(dbmodel.UserWhere.Username.EQ(username)).One(ctx, i.db)
	if err != nil {
		errUsername := errors.New("username not found")
		return model.User{}, errUsername
	}
	if userdb.Password != password {
		errPassword := errors.New("password not match")
		return model.User{}, errPassword
	}
	var user model.User
	user.ID = userdb.ID
	user.Username = userdb.Username
	user.Password = userdb.Password
	user.Email = userdb.Email
	user.Age = int64(user.Age)
	return user, nil
}
