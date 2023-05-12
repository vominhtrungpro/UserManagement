package user

import (
	"encoding/json"
	"net/http"
	"strings"
	"vominhtrungpro/usermanagement/internal/controller/user"
	"vominhtrungpro/usermanagement/internal/httpserver"
)

func (h Handler) CreateUser() http.HandlerFunc {
	return httpserver.HandlerErr(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		input := user.CreateInput{}
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			return &httpserver.Error{Status: http.StatusBadRequest, Code: ErrCodeValidationFailed, Desc: err.Error()}
		}

		if err := validatecreate(input); err != nil {
			return &httpserver.Error{Status: http.StatusBadRequest, Code: ErrCodeValidationFailed, Desc: err.Error()}
		}

		resp, err := h.userCtrl.CreateUser(ctx, input)
		if err != nil {
			return &httpserver.Error{Status: http.StatusBadRequest, Code: ErrCodeValidationFailed, Desc: err.Error()}
		}

		httpserver.RespondJSON(w, resp)
		return nil
	})
}

func validatecreate(input user.CreateInput) error {
	if strings.TrimSpace(input.Username) == "" {
		return errInvalidUsername
	}

	if strings.TrimSpace(input.Password) == "" {
		return errInvalidPassword
	}

	if strings.TrimSpace(input.Email) == "" {
		return errInvalidEmail
	}

	if input.Age <= 0 {
		return errInvalidAge
	}

	return nil
}
