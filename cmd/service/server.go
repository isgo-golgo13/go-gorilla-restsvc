package main

import (
	"log"
	"net/http"

	"github.com/isgo-golgo13/gorilla-restsvc/router"
)

func main() {

	router := router.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
