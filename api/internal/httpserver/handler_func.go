package httpserver

import "net/http"

// HandlerErr handle error & data then return to client
func HandlerErr(h func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if err := h(writer, request); err != nil {
			RespondJSON(writer, err)
		}
	}
}
