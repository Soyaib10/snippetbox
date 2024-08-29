package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

// home shows homepage of the app
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w) // Using notFound() helper
		return
	}

	// Initialize a slice containing the paths to the two files. Note that the home.page.tmpl file must be the *first* file in the slice.
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err) // Using serverError() helper
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err) // Using serverError() helper.
	}
}

// showSnippet shows a specific snippet
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w) // Using notFount() helper
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// createSnippet creates a new snippet
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed) // Using clientError() helper
		return
	}
	w.Write([]byte("Create a new snippet..."))
}
