package user

import (
	"context"
	"errors"
	"vominhtrungpro/usermanagement/internal/model"
	dbmodel "vominhtrungpro/usermanagement/internal/repository/dbmodel"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (i impl) UpdateUser(ctx context.Context, m model.User) (model.User, error) {

	userdb, err := dbmodel.Users(dbmodel.UserWhere.ID.EQ(m.ID)).One(ctx, i.db)
	if err != nil {
		return model.User{}, err
	}

	existingUsername, err := dbmodel.Users(dbmodel.UserWhere.Username.EQ(m.Username), dbmodel.UserWhere.ID.NEQ(m.ID)).Exists(ctx, i.db)
	if err != nil {
		return model.User{}, err
	}
	if existingUsername {
		errNotFound := errors.New("username exist")
		return model.User{}, errNotFound
	}

	existingEmail, err := dbmodel.Users(dbmodel.UserWhere.Email.EQ(m.Email), dbmodel.UserWhere.ID.NEQ(m.ID)).Exists(ctx, i.db)
	if err != nil {
		return model.User{}, err
	}
	if existingEmail {
		errNotFound := errors.New("email exist")
		return model.User{}, errNotFound
	}

	userdb.Username = m.Username
	userdb.Password = m.Password
	userdb.Age = int(m.Age)
	userdb.Email = m.Email

	if _, errr := userdb.Update(ctx, i.db, boil.Infer()); err != nil {
		return model.User{}, errr
	}
	return m, nil
}
