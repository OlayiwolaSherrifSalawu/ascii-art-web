package main

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
	// Parse all templates — base pages and partials
	templates, err := template.ParseGlob("ui/templates/*.tmpl")
	if err != nil {
		log.Fatal(err)
		return
	}
	templates, err = templates.ParseGlob("ui/templates/partials/*.tmpl")
	if err != nil {
		log.Fatal(err)
		return
	}

	cfg := new(config)

	// Create ASCII service (reads font files from the given directory)
	srv := ascii.NewAsciiService("ui/static/fonts")

	// Create handler, injecting the service and templates
	app := handlers.NewAsciiHandler(srv, templates)

	flag.StringVar(&cfg.port, "port", ":8080", "HTTP PORT ADDRESS")
	flag.Parse()

	cfg.infologger = *log.New(os.Stdout, "INFO \t", log.Ldate|log.Ltime)
	cfg.errorlogger = *log.New(os.Stderr, "ERROR \t", log.Ldate|log.Ltime)

	mux := http.NewServeMux()

	// Static assets
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("ui/static"))))

	// Pages
	mux.HandleFunc("GET /", app.ServeHome)
	mux.HandleFunc("GET /generate", app.ServeGenerate)

	// API / actions
	mux.HandleFunc("POST /ascii-art", app.ServerAscii)
	mux.HandleFunc("POST /ascii-art/download", app.DownloadHandler)

	cfg.infologger.Printf("started server at port %s\n", cfg.port)
	err = http.ListenAndServe(cfg.port, mux)
	cfg.errorlogger.Println(err)
}
