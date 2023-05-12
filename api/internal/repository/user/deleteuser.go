package user

import (
	"context"
	"errors"
	dbmodel "vominhtrungpro/usermanagement/internal/repository/dbmodel"

	pkgerrors "github.com/pkg/errors"
)

func (i impl) DeleteUser(ctx context.Context, id int64) error {
	userdb, err := dbmodel.Users(dbmodel.UserWhere.ID.EQ(id)).One(ctx, i.db)
	if err != nil {
		errNotFound := errors.New("username exist")
		return errNotFound
	}
	_, errr := userdb.Delete(ctx, i.db)
	if errr != nil {
		return pkgerrors.WithStack(errr)
	}
	return nil
}
