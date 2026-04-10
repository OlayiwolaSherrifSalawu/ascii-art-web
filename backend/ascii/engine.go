package ascii

import (
	"strings"
)

type Config struct {
	OutputFile  string
	InputString []string
	Banner      string
	Color       string
	ColorText   string
	ColorWord   [][]bool
}

func Render(args *Config, file []string, color, reset string) (string, error) {
	// to avoid memory overhead using a string builder would be better
	var sb strings.Builder

	for i, val := range args.InputString {
		if val == "" {
			sb.WriteString("\n")
			continue
		}
		for j := 1; j <= 8; j++ {
			for k := 0; k < len(val); k++ {
				if !(val[k] >= 32 && val[k] <= 126) {
					return "", INVALID_CHAR_VAl
				}
				vals := (int(val[k]-32) * 9) + j
				if args.ColorWord[i][k] {
					sb.WriteString(color)
					sb.WriteString(file[vals])
					sb.WriteString(reset)
				} else {
					sb.WriteString(file[vals])
				}
			}
			sb.WriteString("\n")
		}
	}

	return sb.String(), nil

}
