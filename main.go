package main

import (
	"embed"
	"law-finder/router"
	"log"
	"net/http"
)

//go:embed static/*.md
var laws embed.FS

const (
	port = "8080"
)

func StartServer() {
	r := router.New(laws)
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r.Routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	StartServer()
}
