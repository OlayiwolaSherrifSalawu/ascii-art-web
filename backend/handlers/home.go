package handlers

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/OlayiwolaSherrifSalawu/ascii-art-web.git/backend/ascii"
)

type Handler struct {
	asciiService ascii.AsciiServiceInt
	templates    *template.Template
}

func NewAsciiHandler(asciiService ascii.AsciiServiceInt, templates *template.Template) *Handler {
	return &Handler{
		asciiService: asciiService,
		templates:    templates,
	}
}

func (h *Handler) ServerAscii(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	Text := r.FormValue("text")
	Banner := r.FormValue("banner")
	if Banner == "" || Text == "" {
		h.clientError(w, http.StatusBadRequest)
		return
	}

	Text = strings.ReplaceAll(Text, "\r\n", "\n")
	Result, err := h.asciiService.GenerateAscii(Text, Banner)
	if err != nil {
		h.serverError(w, err)
		return
	}

	// fmt.Fprintf(w, "%s", Result)
	err = h.templates.ExecuteTemplate(w, "result", Result)
	if err != nil {
		h.serverError(w, err)
		return
	}
}

func (h *Handler) ServeHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		h.notFound(w)
		return
	}
	if r.Method != http.MethodGet {
		h.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	err := h.templates.ExecuteTemplate(w, "base", nil)
	if err != nil {
		h.serverError(w, err)
		return
	}
}
