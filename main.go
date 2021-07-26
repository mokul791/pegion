package main

import (
	"log"
	"net/http"
)

/*
URL Patterns		Handler			Action
/					home			Display the home page
/pegion				showPegion		Display a specific pegion
/pegion/create		createPegion	Create a new pegion
*/

// Define a home handler function which writes a byte slice containing
// "Hello from Pegion" as the response body.
// It satisfies the http.Handler interface.
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Pegion"))
}

func showPegion(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Displaying a specific pegion..."))
}

func createPegion(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new pegion..."))
}

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
