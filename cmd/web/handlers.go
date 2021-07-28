package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

/*
URL Patterns			Handler					Action
/						home					Display the home page
/pegion					showPegion				Display a specific pegion
/pegion/create			createPegion			Create a new pegion
/static/				http.FileServer			Serve a specific static file
*/

// Define a home handler function which writes a byte slice containing
// "Hello from Pegion" as the response body.
// It satisfies the http.Handler interface.
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	// w.Write([]byte("Hello from Pegion"))
}

func showPegion(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Displaying a specific perion with id %d...", id)
	// w.Write([]byte("Displaying a specific pegion..."))
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
