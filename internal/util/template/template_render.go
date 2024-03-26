package template

import (
	"html/template"
	"log"
	"net/http"
)

const basicLocation = "client/web/"
const templateLocation = basicLocation + "template/"
const fileExtension = ".html"

func RenderWithHeader(w http.ResponseWriter, location string, data interface{}) {
	headerLocation := "header"
	tmpl := template.Must(template.ParseFiles(makePath(location, "basic"), makePath(headerLocation, "template")))

	err := tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
	}
}

func Render(w http.ResponseWriter, location string, data interface{}) {
	tmpl := template.Must(template.ParseFiles(makePath(location, "basic")))
	err := tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
	}
}

func makePath(location, docsType string) string {
	var path string

	switch docsType {
	case "basic":
		path = basicLocation + location + fileExtension
		break
	case "template":
		path = templateLocation + location + fileExtension
		break
	}
	return path
}
