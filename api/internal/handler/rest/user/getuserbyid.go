package user

import (
	"net/http"
	"strconv"

	"vominhtrungpro/usermanagement/internal/httpserver"

	"github.com/go-chi/chi/v5"
)

// Create creates new product
func (h Handler) GetUserByID() http.HandlerFunc {
	return httpserver.HandlerErr(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()
		id := chi.URLParam(r, "id")
		i, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			panic(err)
		}
		resp, err := h.userCtrl.GetUserByID(ctx, i)
		if err != nil {
			return err
		}

		httpserver.RespondJSON(w, resp)

		return nil
	})
}
