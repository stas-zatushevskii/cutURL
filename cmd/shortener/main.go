package main

import (
	"cutURL/cmd/shortener/config"
	"cutURL/internal/routers"
	"log"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()
	router := routers.RouterNew(cfg.BaseURL)
	log.Fatal(http.ListenAndServe(cfg.ServerURL, router))
}
