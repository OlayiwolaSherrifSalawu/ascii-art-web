package handlers

import (
	"net/http"

	"github.com/OlayiwolaSherrifSalawu/ascii-art-web.git/backend/ascii"
)

type Handler struct {
	asciiService ascii.AsciiService
}

func NewAsciiHandler(asciiService ascii.AsciiService) *Handler {
	return &Handler{asciiService: asciiService}
}

func (h *Handler) ServerAscii(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) ServeHome(w http.ResponseWriter, r *http.Request) {

}
