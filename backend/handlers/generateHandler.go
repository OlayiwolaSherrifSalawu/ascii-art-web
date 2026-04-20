package handlers

import "net/http"

// ServeGenerate serves the ASCII Generator tool page at /generate.
// It renders the "generate-base" template which contains the full
// generator UI including the HTMX-powered live preview form.
func (h *Handler) ServeGenerate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	err := h.templates.ExecuteTemplate(w, "generate-base", nil)
	if err != nil {
		h.serverError(w, err)
		return
	}
}
