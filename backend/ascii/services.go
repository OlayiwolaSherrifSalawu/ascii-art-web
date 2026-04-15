package ascii

type AsciiService interface {
	Render(args *Config, file []string) (string, error)
}

type asciiService struct {
	fontPath string
	cache    map[string][]string
}

func NewAsciiService(fontPaths string) AsciiService {
	return &asciiService{
		fontPath: fontPaths,
		cache:    make(map[string][]string),
	}
}
