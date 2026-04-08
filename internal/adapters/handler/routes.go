package handler

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

func Catalog(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("templates/catalog.html")
	if err != nil {
		fmt.Println("error parsing templates:", err)
		os.Exit(1)
	}
	err = templ.Execute(w, nil)
	if err != nil {
		fmt.Println("error executing template:", err)
		os.Exit(1)
	}
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("templates/create-post.html")
	if err != nil {
		fmt.Println("error parsing templates:", err)
		os.Exit(1)
	}
	err = templ.Execute(w, nil)
	if err != nil {
		fmt.Println("error executing template:", err)
		os.Exit(1)
	}
}

func Archive(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("templates/archive.html")
	if err != nil {
		fmt.Println("error parsing templates:", err)
		os.Exit(1)
	}
	err = templ.Execute(w, nil)
	if err != nil {
		fmt.Println("error executing template:", err)
		os.Exit(1)
	}
}
