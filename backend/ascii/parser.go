package ascii

import (
	"os"
	"path/filepath"
	"strings"
)

func LoadBanner(args Utils) ([]string, error) {
	var file []byte
	var err error
	hasTxtSuffix := strings.HasSuffix(args.Banner, ".txt")
	if !hasTxtSuffix {
		args.Banner = args.Banner + ".txt"
	}
	file, err = os.ReadFile(filepath.Join("fonts/", args.Banner))
	if err != nil {
		return nil, err
	}
	if len(file) <= 1 {
		return nil, EMPTY_FILE
	}
	return strings.Split(string(file), "\n"), nil
}
