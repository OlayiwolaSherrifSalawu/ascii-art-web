package web

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/OlayiwolaSherrifSalawu/ascii-art-web.git/backend/ascii"
	"github.com/OlayiwolaSherrifSalawu/ascii-art-web.git/backend/handlers"
)

type config struct {
	port        string
	infologger  log.Logger
	errorlogger log.Logger
}

func main() {
	templates, err := template.ParseGlob("ui/templates/**/*.html")
	if err != nil {
		log.Println(err)
		return
	}
	cfg := new(config)
	// create a new services
	srv := ascii.NewAsciiService("ui/static/fonts")
	// create a new app handler using the ascii services that was created
	app := handlers.NewAsciiHandler(srv, templates)

	flag.StringVar(&cfg.port, "port", ":8080", "HTTP PORT ADDRESS")
	flag.Parse()
	// create a new logger to log errors and infos  on the terminal
	cfg.infologger = *log.New(os.Stdout, "INFO \t", log.Ldate|log.Ltime)
	cfg.errorlogger = *log.New(os.Stderr, "ERROR \t", log.Ldate|log.Ltime)

	// NewServerMux for cleaner routing
	mux := http.NewServeMux()
	mux.HandleFunc("POST /ascii-art", app.ServerAscii)
	cfg.infologger.Printf("started server at port %s\n", cfg.port)
	err = http.ListenAndServe(cfg.port, mux)
	cfg.errorlogger.Println(err)
}
