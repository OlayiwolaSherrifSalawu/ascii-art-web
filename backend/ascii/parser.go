package ascii

import (
	"os"
	"path/filepath"
	"strings"
)

func LoadBanner(Banner string) ([]byte, error) {
	var file []byte
	var err error
	if strings.HasSuffix(Banner, ".txt") {
		file, err = os.ReadFile(filepath.Join("fonts/", Banner))
		if err != nil {
			return nil, err
		}
		if len(file) <= 1 {
			return nil, EMPTY_FILE
		}

	} else {
		file, err = os.ReadFile(filepath.Join("fonts/", Banner+".txt"))
		if err != nil {
			return nil, err
		}
		if len(file) <= 1 {
			return nil, EMPTY_FILE
		}
	}
	return file, nil
}
