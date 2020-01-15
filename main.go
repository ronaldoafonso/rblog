package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}

	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)

	fmt.Println("Ready to serve rblog requests ...")
	server.ListenAndServe()
}

var templates []string = []string{
	"layout.html",
	"head.html",
	"header.html",
	"index.html",
	"footer.html",
}

// "index" page for rblog.
func index(w http.ResponseWriter, r *http.Request) {
	templates[3] = "index.html"
	fmt.Println(templates)
	indexTemplate, err := template.ParseFiles(templates...)
	if err == nil {
		indexTemplate.ExecuteTemplate(w, "layout", "")
	} else {
		fmt.Println(err)
	}
}

// "about" page for rblog.
func about(w http.ResponseWriter, r *http.Request) {
	templates[3] = "about.html"
	fmt.Println(templates)
	aboutTemplate, err := template.ParseFiles(templates...)
	if err == nil {
		aboutTemplate.ExecuteTemplate(w, "layout", "")
	} else {
		fmt.Println(err)
	}
}
