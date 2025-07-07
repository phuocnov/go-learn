package main

import (
	"github.com/phuocnov/golang-webserver/pkg/models"
)

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
