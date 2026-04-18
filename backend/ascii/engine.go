package ascii

import (
	"strings"
)

type config struct {
	InputString []string
	Color       string
	Reset       string
	ColorWord   [][]bool
}

func (a asciiService) Render(args *config, file []string) (string, error) {
	// to avoid memory overhead using a string builder would be better
	var sb strings.Builder
	if args.ColorWord == nil {
		return "", NOCOLORWORD
	}
	for i, val := range args.InputString {
		if val == "" {
			sb.WriteString("\n")
			continue
		}
		for j := 1; j <= 8; j++ {
			for k, ch := range val {
				if !(ch >= 32 && ch <= 126) {
					return "", INVALID_CHAR_VAl
				}
				vals := (int(ch-32) * 9) + j
				if i < len(args.ColorWord) && k < len(args.ColorWord[i]) && args.ColorWord[i][k] {
					sb.WriteString(args.Color)
					sb.WriteString(file[vals])
					sb.WriteString(args.Reset)
				} else {
					sb.WriteString(file[vals])
				}
			}
			sb.WriteString("\n")
		}
	}

	return sb.String(), nil
}
