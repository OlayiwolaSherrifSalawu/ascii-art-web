package ascii

import (
	"log"
	"os"
	"strings"

	"github.com/OlayiwolaSherrifSalawu/ascii-art-web.git/backend/ascii/utils"
)

type AsciiServiceInt interface {
	GenerateAscii(text, banner string) (string, error)
	ValidBanner(name string) bool
}

type asciiService struct {
	fontPath    string
	cache       map[string][]string
	validBanner map[string]bool
}

func NewAsciiService(fontPaths string) AsciiServiceInt {
	valid := make(map[string]bool)
	banners, err := os.ReadDir(fontPaths)
	if err != nil {
		log.Fatal("could not read directory", err)
	}
	for _, banner := range banners {
		if strings.HasSuffix(banner.Name(), ".txt") {
			name := strings.TrimSuffix(banner.Name(), ".txt")
			valid[name] = true
		}
	}
	return &asciiService{
		fontPath:    fontPaths,
		cache:       make(map[string][]string),
		validBanner: valid,
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
	cfg := &config{
		Color:       "",
		ColorWord:   colorWord,
		InputString: inputSlice,
		Reset:       "",
	}
	return a.Render(cfg, bannerSlices)
}

