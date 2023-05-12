package user

import (
	"net/http"

	"vominhtrungpro/usermanagement/internal/httpserver"
)

const ErrCodeValidationFailed = "validation_failed"

var (
	errInvalidUsername = &httpserver.Error{Status: http.StatusBadRequest, Code: ErrCodeValidationFailed, Desc: "Invalid User Username"}
	errInvalidPassword = &httpserver.Error{Status: http.StatusBadRequest, Code: ErrCodeValidationFailed, Desc: "Invalid User Password"}
	errInvalidEmail    = &httpserver.Error{Status: http.StatusBadRequest, Code: ErrCodeValidationFailed, Desc: "Invalid User Email"}
	errInvalidAge      = &httpserver.Error{Status: http.StatusBadRequest, Code: ErrCodeValidationFailed, Desc: "Invalid User Age"}
)
