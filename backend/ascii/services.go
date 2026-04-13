package ascii

type AsciiService interface {
	Render(args *Config, file []string) (string, error)
	LoadBanner(args Utils) ([]string, error)
}

type asciiService struct {
	fontPath string
}

func NewAsciiService(fontPaths string) AsciiService {
	return &asciiService{fontPath: fontPaths}
}
