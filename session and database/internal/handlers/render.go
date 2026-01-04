package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var TemplateCache map[string]*template.Template

func InitTemplate() error {
	TemplateCache = make(map[string]*template.Template)

	// Get all page templates
	pages, err := filepath.Glob("web/templates/pages/*.html")
	if err != nil {
		return err
	}

	for _, page := range pages {
		// Parse base layout + partials + the specific page
		name := filepath.Base(page)

		ts, err := template.ParseFiles(
			"web/templates/layouts/base.html",
			page,
		)

		if err != nil {
			return err
		}

		// Add all partials
		ts, err = ts.ParseGlob("web/templates/partials/*.html")
		if err != nil {
			return err
		}

		TemplateCache[name] = ts
	}
	return nil
}

func Render(w http.ResponseWriter, tmpl string, data interface{}) {
	ts, ok := TemplateCache[tmpl]

	if !ok {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	err := ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
