package handlers

import (
	"errors"
	"html/template"
	"net/http"

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
	Result, err := h.generateArt(r)
	if errors.Is(err, BAD_REQUEST) {
		h.clientError(w, http.StatusBadRequest)
		return
	}
	if errors.Is(err, INVALID_CHAR) {
		h.templates.ExecuteTemplate(w, "error", err)
		return
	}
	if errors.Is(err, EMPTY_STRING) {
		h.templates.ExecuteTemplate(w, "result", "")
		return	
	}
	if err !=nil{

	}
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
