package handlers

import (
	"log"
	"net/http"

	"github.com/OlayiwolaSherrifSalawu/ascii-art-web.git/backend/ascii"
)

type Handler struct {
	asciiService ascii.AsciiServiceInt
	errorLogger  log.Logger
	infoLogger   log.Logger
}

func NewAsciiHandler(asciiService ascii.AsciiServiceInt) *Handler {
	return &Handler{asciiService: asciiService}
}

func (h *Handler) ServerAscii(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) ServeHome(w http.ResponseWriter, r *http.Request) {

}
