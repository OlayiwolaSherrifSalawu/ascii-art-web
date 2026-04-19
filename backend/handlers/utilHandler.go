package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/OlayiwolaSherrifSalawu/ascii-art-web.git/backend/ascii"
)

func (h *Handler) generateArt(r *http.Request) (string, error) {
	Text := r.FormValue("text")
	Banner := r.FormValue("banner")
	if Banner == "" {
		return "", BAD_REQUEST
	}
	isValidBanner := h.asciiService.ValidBanner(Banner)
	if !isValidBanner {
		return "", INVALID_BANNER
	}
	Text = strings.ReplaceAll(Text, "\r\n", "\n")
	if Text == "" {
		return "", EMPTY_STRING
	}
	Result, err := h.asciiService.GenerateAscii(Text, Banner)
	if err != nil {
		if errors.Is(err, ascii.INVALID_CHAR_VAl) {
			return "", INVALID_CHAR
		}
		return "", err
	}
	return Result, nil
}
