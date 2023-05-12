package user

import (
	"net/http"
	"strconv"
	"vominhtrungpro/usermanagement/internal/httpserver"

	"github.com/go-chi/chi/v5"
)

func (h Handler) DeleteUser() http.HandlerFunc {
	return httpserver.HandlerErr(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()
		id := chi.URLParam(r, "id")
		i, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			panic(err)
		}
		errr := h.userCtrl.DeleteUser(ctx, i)
		if errr != nil {
			return &httpserver.Error{Status: http.StatusBadRequest, Code: ErrCodeValidationFailed, Desc: err.Error()}
		}
		httpserver.RespondJSON(w, "Success")
		return nil
	})
}
