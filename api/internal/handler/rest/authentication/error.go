package authentication

import (
	"net/http"
	"vominhtrungpro/usermanagement/internal/httpserver"
)

const ErrCodeValidationFailed = "validation_failed"

var (
	errInvalidUsername = &httpserver.Error{Status: http.StatusBadRequest, Code: ErrCodeValidationFailed, Desc: "Invalid Username"}
	errInvalidPassword = &httpserver.Error{Status: http.StatusBadRequest, Code: ErrCodeValidationFailed, Desc: "Invalid Password"}
)
