package handlers

import "net/http"

func (h *Handler) clientError(w http.ResponseWriter, err int) {
	http.Error(w, http.StatusText(err), err)
}
