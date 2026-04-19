package ascii

import (
	"os"
	"path/filepath"
	"strings"
)

func (a asciiService) LoadBanner(Banner string) ([]string, error) {
	var file []byte
	var err error

	hasTxtSuffix := strings.HasSuffix(Banner, ".txt")
	if !hasTxtSuffix {
		Banner = Banner + ".txt"
	}
	if lines, ok := a.cache[Banner]; ok {
		return lines, nil
	}
	file, err = os.ReadFile(filepath.Join(a.fontPath, Banner))
	if err != nil {
		return nil, err
	}
	if len(file) <= 1 {
		return nil, EMPTY_FILE
	}
	a.cache[Banner] = strings.Split(string(file), "\n")
	return strings.Split(string(file), "\n"), nil
}

func (a asciiService) ValidBanner(name string) bool {
	return a.validBanner[name]
}
