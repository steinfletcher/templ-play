package main

import (
	"log"
	"net/http"

	"github.com/steinfletcher/templ-play/web"
)

//go:generate templ generate

func main() {
	log.Fatal(http.ListenAndServe(":8000", web.NewRouter()))
}
