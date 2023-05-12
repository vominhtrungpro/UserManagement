package user

import "vominhtrungpro/usermanagement/internal/controller/user"

type Handler struct {
	userCtrl user.Controller
}

func New(userCtrl user.Controller) Handler {
	return Handler{
		userCtrl: userCtrl,
	}
}
