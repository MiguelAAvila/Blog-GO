package main

import (
	"net/url"

	"mavila_frosado.net/test1/pkg/models"
)

type templateData struct {
	Blogs          []*models.Blog
	ErrorsFromForm map[string]string
	FormData       url.Values
}
