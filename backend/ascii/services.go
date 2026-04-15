package ascii

import (
	"strings"

	"github.com/OlayiwolaSherrifSalawu/ascii-art-web.git/backend/ascii/utils"
)

type AsciiServiceInt interface {
	GenerateAscii(text, banner string) (string, error)
}

type asciiService struct {
	fontPath string
	cache    map[string][]string
}

func NewAsciiService(fontPaths string) AsciiServiceInt {
	return &asciiService{
		fontPath: fontPaths,
		cache:    make(map[string][]string),
	}
}

func (a *asciiService) GenerateAscii(text, banner string) (string, error) {
	if text == "" || banner == "" {
		return "", INVALID_INPUTS
	}
	inputSlice := strings.Split(text, "\n")
	bannerSlices, err := a.LoadBanner(banner)

	if err != nil {
		return "", CANT_READ_BANNER
	}
	colorWord := utils.GetIndex(inputSlice, "")
	cfg := &Config{
		Color:       "",
		ColorWord:   colorWord,
		InputString: inputSlice,
		Reset:       "",
	}
	return a.Render(cfg, bannerSlices)
}
