package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func NoughtsAndCrosses(w http.ResponseWriter, r *http.Request) {
	//if r.URL.Path != "/noughtsAndCrosses" {
	//	http.NotFound(w, r)
	//	return
	//}

	templateFiles := []string{
		"./frontend/noughtsAndCrosses/page.tmpl",
		"./frontend/base/base.tmpl",
	}

	ts, err := template.ParseFiles(templateFiles...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}