package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/Soyaib10/snippetbox/pkg/models"
)

// home shows homepage of the app
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w) // Using notFound() helper
		return
	}

	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Create an instance of a templateData struct holding the slice of snippets
	data := &templateData{Snippets: s}


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

	err = ts.Execute(w, data)
	if err != nil {
		app.serverError(w, err) // Using serverError() helper.
	}
}

// showSnippet shows a specific snippet
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	s, err := app.snippets.Get(id)
	if err == models.ErrNoRecord {
		app.notFound(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	// Create an instance of a templateData struct holding the snippet data.
	data := &templateData{Snippet: s}

	// Initialize a slice containing the paths to the show.page.tmpl file, plus the base layout and footer partial that we made earlier.
	files := []string{
		"./ui/html/show.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	// Parse the template files...
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// And then execute them. Notice how we are passing in the snippet data (a models.Snippet struct) as the final parameter.
	err = ts.Execute(w, data)
	if err != nil {
		app.serverError(w, err)
	}
}

// createSnippet creates a new snippet
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed) // Using clientError() helper
		return
	}

	// Create some variables holding dummy data. We'll remove these later on during the build.
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi"
	expires := "7"

	// Pass the data to the snippetModel.Insert() method, receiving the ID of hte new record.
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Redirect the user to the relevant page for the snippet.
	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}
