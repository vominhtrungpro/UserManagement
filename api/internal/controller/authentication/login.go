package authentication

import (
	"context"
	"errors"
)

type LoginInput struct {
	Username string
	Password string
}

func (i impl) Login(ctx context.Context, input LoginInput) (TokenOutput, error) {
	var token TokenOutput
	userdb, err := i.repo.User().GetUserByUsername(ctx, input.Username)
	if err != nil {
		return token, err
	}

	if userdb.Password != input.Password {
		return token, errors.New("password not match")
	}

	accessToken, errr := CreateAccessToken(userdb)
	if errr != nil {
		return token, errr
	}
	refreshToken, errrr := CreateRefreshToken()
	if errrr != nil {
		return token, errrr
	}
	updateErr := i.repo.User().UpdateRefreshToken(ctx, userdb, refreshToken)
	if updateErr != nil {
		return token, updateErr
	}
	token.AccessToken = accessToken
	token.RefreshToken = refreshToken
	return token, nil
}
