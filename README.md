# TypAscii

> Transform any text into bold ASCII art instantly — powered by Go and HTMX.

**[Live Demo →](https://ascii-art-web-4suu.onrender.com)**

```
 _              _   _
| |            | | | |
| |__     ___  | | | |   ___
|  _ \   / _ \ | | | |  / _ \
| | | | |  __/ | | | | | (_) |
|_| |_|  \___| |_| |_|  \___/
```

## Overview

TypAscii is a fast, lightweight web application that converts plain text into ASCII art. Type anything, pick a font style, and get your output live — no page reloads, no friction. Copy it to your clipboard or download it as a `.txt` file.

## Features

- **Live Preview** — Output updates as you type, powered by HTMX with zero page reloads
- **Multiple Font Styles** — Ship with Standard, Shadow, and Thinker-Toy fonts, easily extensible
- **One-Click Copy** — Copy raw ASCII output straight to your clipboard
- **Export as .txt** — Download your creation as a plain text file
- **Dark / Light Theme** — Toggle between themes with persistent preference via `localStorage`
- **Responsive UI** — Clean, minimal interface that works across devices

## Tech Stack

| Layer | Technology |
|-------|-----------|
| Backend | Go (net/http) |
| Frontend | HTML Templates + HTMX |
| Styling | Custom CSS (no framework) |
| Fonts | Plain `.txt` banner files |

## Project Structure

```
.
├── cmd/web/main.go              # Application entry point
├── backend/
│   ├── ascii/                   # Core ASCII generation engine
│   │   ├── engine.go            # Rendering logic
│   │   ├── parser.go            # Banner file loading & caching
│   │   ├── services.go          # Service interface & orchestration
│   │   ├── color.go             # Color support utilities
│   │   └── errors.go            # Typed error constants
│   ├── handlers/                # HTTP request handlers
│   │   ├── coreHandlers.go      # Home page & ASCII generation
│   │   ├── generateHandler.go   # Generator app page
│   │   ├── downloadHandler.go   # .txt file download
│   │   └── errors.go            # Handler error helpers
│   ├── models/                  # Data models
│   ├── database/                # Database layer
│   └── server/middleware.go     # Server middleware
├── ui/
│   ├── templates/               # Go HTML templates
│   │   ├── base.tmpl            # Shared layout (nav, footer, theme)
│   │   ├── home.tmpl            # Landing page
│   │   ├── generate.tmpl        # Generator app (HTMX form + output)
│   │   └── partials/            # Reusable template fragments
│   └── static/
│       ├── css/                 # Stylesheets
│       └── fonts/               # ASCII banner font files (.txt)
├── go.mod
└── README.md
```

## Getting Started

### Prerequisites

- **Go 1.22+**

### Run Locally

```bash
# Clone the repository
git clone https://github.com/OlayiwolaSherrifSalawu/ascii-art-web.git
cd ascii-art-web

# Run the server (defaults to port 8080)
go run ./cmd/web

# Or specify a custom port
go run ./cmd/web -port :3000
```

Open [http://localhost:8080](http://localhost:8080) in your browser.

## API Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/` | Landing page |
| `GET` | `/generate` | ASCII art generator app |
| `POST` | `/ascii-art` | Generate ASCII art (returns HTML fragment) |
| `POST` | `/ascii-art/download` | Download generated art as `.txt` |

## Adding Custom Fonts

Drop a new `.txt` banner file into `ui/static/fonts/`. The file must follow the standard 8-line-per-character ASCII banner format (characters 32–126, 9 lines each including blank separators). The font name is derived from the filename — e.g., `doom.txt` becomes available as the `doem` font option.

## Architecture

The application follows a clean separation of concerns:

- **`ascii` package** — Pure generation logic. Parses banner files, caches them in memory, and renders text to ASCII output. No HTTP awareness.
- **`handlers` package** — HTTP layer. Receives requests, validates input, calls the ASCII service, and executes templates.
- **Templates** — HTMX-driven frontend. The generator form fires `hx-post` on input change, and the server returns rendered HTML fragments — no JavaScript framework needed.

## License

MIT

---

Built by [Olayiwola Sherrif Salawu](https://github.com/OlayiwolaSherrifSalawu)
