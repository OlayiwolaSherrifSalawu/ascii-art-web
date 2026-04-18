package handlers

import (
	"errors"
	"net/http"
)

func (h *Handler) DownloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art/download" {
		h.notFound(w)
		return
	}
	Result, err := h.generateArt(r)
	if errors.Is(err, INVALID_CHAR) {
		h.templates.ExecuteTemplate(w, "error", err)
		h.clientError(w, http.StatusBadRequest)
		return
	}
	fileName := "ascii-art"
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", "application/txt")
	w.Write([]byte(Result))
	http.ServeFile(w, r, fileName)
}
