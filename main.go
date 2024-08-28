package main

import (
	"log"
	"net/http"

	"github.com/powderbluecrayon/memory-game/templates"
	"github.com/powderbluecrayon/memory-game/templates/layouts"

	"github.com/a-h/templ"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.Handle("/", templ.Handler(layouts.GameLayout()))
	http.Handle("/flip", templ.Handler(templates.Flip(1, 1)))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
