package main

import (
	"cutURL/internal/routers"
	"log"
	"net/http"
)

func main() {
	router := routers.RouterNew()
	log.Fatal(http.ListenAndServe(":8080", router))
}
