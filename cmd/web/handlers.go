package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"swarmtree.com/pegion/pkg/models"
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
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
	}

	// w.Write([]byte("Hello from Pegion"))
}

func (app *application) showPegion(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	s, err := app.pegions.Get(id)
	if err == models.ErrNoRecord {
		app.notFound(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}
	fmt.Fprintf(w, "Displaying your pegion with id %d...\n", id)
	fmt.Fprintf(w, "%v", s)

	// w.Write([]byte("Displaying a specific pegion..."))
}

func (app *application) createPegion(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		// w.WriteHeader(405)
		// w.Write([]byte("Method not Allowed"))
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "O Snail"
	content := "Hello Snail!\nDo you feel lonely"
	expires := "7"

	id, err := app.pegions.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/pegion?id=%d", id), http.StatusSeeOther)

	w.Write([]byte("Create a new pegion..."))
}
