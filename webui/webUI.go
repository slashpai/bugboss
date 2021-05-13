package webui

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

// Data ...
type Data struct {
	Header     []string
	DataFields [][]string
}

// ShowUI ...
func ShowUI(data [][]string, header []string) {
	fp := path.Join("templates", "index.html")
	tmpl := template.Must(template.ParseFiles(fp))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, Data{Header: header, DataFields: data})
	})

	log.Print("Open http://localhost:8090 to view output in browser")
	http.ListenAndServe(":8090", nil)
}
