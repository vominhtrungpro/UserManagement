package authentication

import (
	"encoding/json"
	"net/http"
	"strings"
	"vominhtrungpro/usermanagement/internal/controller/authentication"
	"vominhtrungpro/usermanagement/internal/httpserver"
)

func (h Handler) Login() http.HandlerFunc {
	return httpserver.HandlerErr(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		input := authentication.LoginInput{}
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			return err
		}

		if err := validatelogin(input); err != nil {
			return err
		}

		tokenresp, err := h.authenCtrl.Login(ctx, input)
		if err != nil {
			return &httpserver.Error{Status: http.StatusBadRequest, Code: ErrCodeValidationFailed, Desc: err.Error()}
		}

		httpserver.RespondJSON(w, tokenresp)
		return nil
	})
}

func validatelogin(input authentication.LoginInput) error {
	if strings.TrimSpace(input.Username) == "" {
		return errInvalidUsername
	}

	if strings.TrimSpace(input.Password) == "" {
		return errInvalidPassword
	}
	return nil
}
