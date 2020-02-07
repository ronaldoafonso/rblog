package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"os"
)

var (
	templates []string = []string{
		"layout.html",
		"head.html",
		"header.html",
		"index.html",
		"footer.html",
	}
	articles map[string]func(http.ResponseWriter, *http.Request)
)

func main() {
	articles = make(map[string]func(http.ResponseWriter, *http.Request))
	articles["index"] = index
	articles["about"] = about

	router := mux.NewRouter()
	router.HandleFunc("/{article}", handleArticle)
	router.HandleFunc("/", index)
	logRouter := handlers.LoggingHandler(os.Stdout, router)

	server := &http.Server{
		Handler: logRouter,
		Addr:    ":8080",
	}

	fmt.Println("Ready to serve rblog requests ...")
	log.Fatal(server.ListenAndServe())
}

// handleArticle ... Handle an especific article request.
func handleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if article, ok := articles[vars["article"]]; ok {
		article(w, r)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Page Not Found.")
	}
}

// index ... Handle "index" page.
func index(w http.ResponseWriter, r *http.Request) {
	templates[3] = "index.html"
	indexTemplate, err := template.ParseFiles(templates...)
	if err == nil {
		indexTemplate.ExecuteTemplate(w, "layout", "")
	} else {
		fmt.Println(err)
	}
}

// about .. Handle "about" page.
func about(w http.ResponseWriter, r *http.Request) {
	templates[3] = "about.html"
	aboutTemplate, err := template.ParseFiles(templates...)
	if err == nil {
		aboutTemplate.ExecuteTemplate(w, "layout", "")
	} else {
		fmt.Println(err)
	}
}
