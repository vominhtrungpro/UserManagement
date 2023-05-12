package user

import (
	"context"
	"errors"
	"vominhtrungpro/usermanagement/internal/model"
	dbmodel "vominhtrungpro/usermanagement/internal/repository/dbmodel"
	"vominhtrungpro/usermanagement/internal/repository/generator"

	pkgerrors "github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (i impl) CreateUser(ctx context.Context, m model.User) (model.User, error) {
	id, err := generator.ProductSNF.Generate()
	if err != nil {
		return model.User{}, err
	}

	userdb, _ := dbmodel.Users(dbmodel.UserWhere.Username.EQ(m.Username)).One(ctx, i.db)
	if userdb != nil {
		errNotFound := errors.New("username exist")
		return model.User{}, errNotFound
	}

	userdb2, _ := dbmodel.Users(dbmodel.UserWhere.Email.EQ(m.Email)).One(ctx, i.db)
	if userdb2 != nil {
		errNotFound := errors.New("email exist")
		return model.User{}, errNotFound
	}

	o := dbmodel.User{
		ID:       id,
		Username: m.Username,
		Password: m.Password,
		Email:    m.Email,
		Age:      int(m.Age),
	}

	if err := o.Insert(ctx, i.db, boil.Infer()); err != nil {
		return model.User{}, pkgerrors.WithStack(err)
	}

	m.ID = id
	return m, nil
}
