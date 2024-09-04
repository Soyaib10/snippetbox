package main

import "github.com/Soyaib10/snippetbox/pkg/models"

// templateData acts as the holding structure for any dynamic data that we want to pass to our HTML templates as we know s html/template package allow you to pass in only one item of dynamic data when rendating a template

type templateData struct {
	Snippet *models.Snippet
}