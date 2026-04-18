package handlers

import (
	"log"
	"net/http"
)

type constError string

func (e constError) Error() string {
	return string(e)
}

const (
	BAD_REQUEST  = constError("BAD REQUEST")
	INVALID_CHAR = constError("INVALID CHARACTER DETECTED")
	EMPTY_STRING = constError("EMPTY STRING")
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
