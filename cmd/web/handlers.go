package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"

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
	app.render(w, r, "create.page.tmpl", nil)
}

// createSnippet creates a new snippet
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	// First we call r.ParseForm() which adds any data in POST request bodies to the r.PostForm map.
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	expires := r.PostForm.Get("expires")

	// Initialize a map to hold any validation errors.
	errors := make(map[string]string)

	// Checking errors
	if strings.TrimSpace(title) == "" {
		errors["title"] = "Title field can't be black"
	} else if utf8.RuneCountInString(title) > 100 {
		errors["title"] = "Title field is too long (max 100 characters allowed)"
	}

	if strings.TrimSpace(content) == "" {
		errors["content"] = "Content field can't be blank"
	}

	if strings.TrimSpace(expires) == "" {
		errors["expires"] = "Expires field can't be blank"
	} else if expires != "365" && expires != "7" && expires != "1" {
		errors["expires"] = "Expires field is invalid"
	}

	// If there are any errors, dump them in a plain text HTTP response and ret from the handler.
	if len(errors) > 0 {
		fmt.Fprint(w, errors)
		return
	}

	// Create a new snippet record in the database using the form data.
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}
