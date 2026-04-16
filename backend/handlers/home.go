package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/OlayiwolaSherrifSalawu/ascii-art-web.git/backend/ascii"
)

type Handler struct {
	asciiService ascii.AsciiServiceInt
	Banner       string
	Text         string
	Result       string
}

func NewAsciiHandler(asciiService ascii.AsciiServiceInt) *Handler {
	return &Handler{asciiService: asciiService}
}

func (h *Handler) ServerAscii(w http.ResponseWriter, r *http.Request) {
	var err error
	if r.Method != http.MethodPost {
		h.clientError(w, http.StatusBadRequest)
		return
	}
	h.Text = r.FormValue("text")
	h.Banner = r.FormValue("banner")
	if h.Banner == "" || h.Text == "" {
		h.clientError(w, http.StatusBadRequest)
		return
	}
	services := ascii.NewAsciiService("ui/static/fonts")
	h.Text = strings.ReplaceAll(h.Text, "\r\n", "\n")
	h.Result, err = services.GenerateAscii(h.Result, h.Banner)
	if err != nil {
		h.serverError(w, err)
		return
	}
	w.Header().Set("Content-type", "text/html")
	fmt.Fprintf(w, "%s", h.Result)
}

func (h *Handler) ServeHome(w http.ResponseWriter, r *http.Request) {

}
