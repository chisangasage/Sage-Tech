package web

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

//IndexHandler main page handler
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		err := tmpl.ExecuteTemplate(w, "index.gohtml", nil)
		if err != nil {
			log.Fatal(err)
		}
	}
}
