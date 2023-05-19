package authentication

import "vominhtrungpro/usermanagement/internal/controller/authentication"

type Handler struct {
	authenCtrl authentication.Controller
}

func New(authenCtrl authentication.Controller) Handler {
	return Handler{
		authenCtrl: authenCtrl,
	}
}
