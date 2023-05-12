package user

import (
	"context"
	"vominhtrungpro/usermanagement/internal/model"
	dbmodel "vominhtrungpro/usermanagement/internal/repository/dbmodel"
)

func (i impl) GetUserByID(ctx context.Context, id int64) (model.User, error) {
	userdb, err := dbmodel.Users(dbmodel.UserWhere.ID.EQ(id)).One(ctx, i.db)
	if err != nil {
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
