package user

import (
	"net/http"

	"vominhtrungpro/usermanagement/internal/httpserver"
)

// Create creates new product
func (h Handler) GetAllUser() http.HandlerFunc {
	return httpserver.HandlerErr(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		resp, err := h.userCtrl.GetAllUser(ctx)
		if err != nil {
			return err
		}

		httpserver.RespondJSON(w, resp)

		return nil
	})
}
