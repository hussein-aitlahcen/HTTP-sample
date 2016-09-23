package main

import (
	"log"
	"net/http"

	"../myserver"
)

const addr = "localhost:12345"

func main() {
	mux := http.NewServeMux()

	mux.Handle("/favicon.ico", http.NotFoundHandler())
	mux.Handle("/visitors", &myserver.MyHandler{})
	mux.Handle("/", &myserver.WelcomeHandler{Message: "TP inte"})

	log.Printf("Now listening on %s...\n", addr)

	server := http.Server{Handler: mux, Addr: addr}

	log.Fatal(server.ListenAndServe())
}
