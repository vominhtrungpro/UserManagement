package user

import "context"

func (i impl) DeleteUser(ctx context.Context, id int64) error {
	return i.repo.User().DeleteUser(ctx, id)
}
