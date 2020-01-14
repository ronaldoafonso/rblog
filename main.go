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

// "index" page for rblog.
func index(w http.ResponseWriter, r *http.Request) {
	indexTemplate, err := template.ParseFiles("index.html", "head.html", "header.html", "footer.html")
	if err == nil {
		indexTemplate.Execute(w, "")
	} else {
		fmt.Println(err)
	}
}

// "about" page for rblog.
func about(w http.ResponseWriter, r *http.Request) {
	aboutTemplate, err := template.ParseFiles("about.html", "head.html", "header.html", "footer.html")
	if err == nil {
		aboutTemplate.Execute(w, "")
	} else {
		fmt.Println(err)
	}
}
