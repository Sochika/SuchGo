package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/Sochika/SuchGo/pkg/config"
)

var functions = template.FuncMap{}

var app *config.SystemConfig

func RefreshTemplates(a *config.SystemConfig) {
	app = a
}

// RenderTemplate renders templates using html
func RenderTemplate(w http.ResponseWriter, tmpl string) {

	pagesCached := app.TemplateCache

	page, ok := pagesCached[tmpl]
	if !ok {
		log.Fatalf("%s Page not Found in Template", tmpl)
	}

	buf := new(bytes.Buffer)

	_ = page.Execute(buf, nil)

	_, err := buf.WriteTo(w)

	if err != nil {
		log.Println("Error writing template to Browser", err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./template/*.page.html")
	if err != nil {
		return myCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		log.Println("Page is currently", page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./template/layout/*.layout.html")

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./template/layout/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
