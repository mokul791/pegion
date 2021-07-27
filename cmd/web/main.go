package main

import (
	"log"
	"net/http"
)

func main() {
	// Initialize a new servermux, router and then register the home function as the
	// handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/pegion", showPegion)
	mux.HandleFunc("/pegion/create", createPegion)

	log.Println("Starting server on :4000")
	// http.ListenAndServe() function starts a new web server.
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
