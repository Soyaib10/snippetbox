package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/Soyaib10/snippetbox/pkg/forms"
	"github.com/Soyaib10/snippetbox/pkg/models"
)

// templateData acts as the holding struct for any dynamic data that we want to pass to our HTML templates as we know s html/template package allow you to pass in only one item of dynamic data when rendating a template
type templateData struct {
	AuthenticatedUser int
	CurrentYear       int
	Flash             string
	Form              *forms.Form
	Snippet           *models.Snippet
	Snippets          []*models.Snippet
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 03:04 PM")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl")) // This essentially gives us a slice of all the 'page' templates for the application
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page) // Base returns the last element of the path

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl")) //  ParseGlob method to add any 'layout' templates to the template set
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
