package user

import (
	"context"
	"database/sql"
	"errors"
	"vominhtrungpro/usermanagement/internal/model"
	dbmodel "vominhtrungpro/usermanagement/internal/repository/dbmodel"
)

func (i impl) GetUserByUsername(ctx context.Context, username string) (model.User, error) {
	userdb, err := dbmodel.Users(dbmodel.UserWhere.Username.EQ(username)).One(ctx, i.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, errors.New("user not found")
		}
		return model.User{}, err
	}
	var user model.User
	user.ID = userdb.ID
	user.Username = userdb.Username
	user.Password = userdb.Password
	user.Email = userdb.Email
	user.Age = int64(user.Age)
	return user, nil
}
