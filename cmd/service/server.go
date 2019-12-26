package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/isgo-golgo13/gorilla-restsvc/router"
)

func main() {

	router := router.NewRouter()
	//log.Fatal(http.ListenAndServe(":8080", router))

	//Start Server
	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		errs <- http.ListenAndServe(":8080", router)
	}()

	log.Print("Exit ", <-errs)

}
