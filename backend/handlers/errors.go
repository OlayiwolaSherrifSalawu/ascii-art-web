package handlers

import (
	"log"
	"net/http"
)

func (h *Handler) clientError(w http.ResponseWriter, err int) {
	http.Error(w, http.StatusText(err), err)
}
func (h *Handler) serverError(w http.ResponseWriter, err error) {
	log.Println(err)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (h *Handler) notFound(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}
