package main

import (
	"log"
	"memorygame/templates"
	"memorygame/templates/layouts"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.Handle("/", templ.Handler(layouts.GameLayout()))
	http.Handle("/flip", templ.Handler(templates.Flip(1, 1)))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
