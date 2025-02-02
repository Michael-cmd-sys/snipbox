package main 

import (
	"html/template"
	"path/filepath"
	"time"
	"io/fs"

	"github.com/Michael-cmd-sys/snipbox/internal/models"
	"github.com/Michael-cmd-sys/snipbox/ui"
)

type templateData struct {
	Snippet *models.Snippet
	Snippets []*models.Snippet
	CurrentYear int
	Form any
	Flash string
	IsAuthenticated bool
	CSRFToken string
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:03")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := fs.Glob(ui.Files, "html/pages/*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		patterns := []string {
			"html/base.tmpl.html",
			"html/partials/*html",
			page,
		}

		ts, err := template.New(name).Funcs(functions).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}
	
	return cache, nil
}

