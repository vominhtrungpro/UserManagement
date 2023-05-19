package user

import (
	"context"
	"time"
	"vominhtrungpro/usermanagement/internal/model"
	dbmodel "vominhtrungpro/usermanagement/internal/repository/dbmodel"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (i impl) UpdateRefreshToken(ctx context.Context, m model.User, token string) error {

	userdb, err := dbmodel.Users(dbmodel.UserWhere.ID.EQ(m.ID)).One(ctx, i.db)
	if err != nil {
		return err
	}

	userdb.Refreshtoken = null.String{String: token, Valid: true}
	expdate := time.Now()
	userdb.RefreshtokenExpiretime = null.Time{Time: expdate, Valid: true}

	if _, errr := userdb.Update(ctx, i.db, boil.Infer()); err != nil {
		return errr
	}
	return nil
}
