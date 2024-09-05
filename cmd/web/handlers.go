package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Soyaib10/snippetbox/pkg/models"
)

// home shows homepage of the app
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Because Pat matches the "/" path exactly, we can now remove the manual of r.URL.Path != "/" from this handler.
	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	// User render helper
	app.render(w, r, "home.page.tmpl", &templateData{
		Snippets: s,
	})

}

// showSnippet shows a specific snippet
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	// Pat doesn't strip the colon from the named capture key, so we need to get the value of ":id" from the query string instead of "id".
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
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

	// Using render helper
	app.render(w, r, "show.page.tmpl", &templateData{
		Snippet: s,
	})
}

// createSnippetForm handler, which for returns a placeholder
func (app *application) createSnippetForm(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet...."))
}

// createSnippet creates a new snippet
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
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
	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}
