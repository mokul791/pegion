package main

import (
	"fmt"
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
	// w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
		// w.Header().Set("Allow", "POST")
		// w.WriteHeader(405)
		// w.Write([]byte("Method not Allowed"))
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new pegion..."))
}
