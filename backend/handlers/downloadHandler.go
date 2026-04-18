package handlers

import (
	"errors"
	"log"
	"net/http"
)

func (h *Handler) DownloadHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("DownloadHandler called")
	Result, err := h.generateArt(r)
	if errors.Is(err, INVALID_CHAR) {
		h.clientError(w, http.StatusBadRequest)
		return
	}
	if err != nil {
		h.clientError(w, http.StatusBadRequest)
		return
	}
	fileName := "ascii-art.txt"
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	w.Write([]byte(Result))
}
