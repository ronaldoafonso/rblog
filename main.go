package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/ronaldoafonso/rblog/rarticle"
	"log"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", rarticle.HandleArticle)
	router.HandleFunc("/{article}", rarticle.HandleArticle)
	logRouter := handlers.LoggingHandler(os.Stdout, router)

	server := &http.Server{
		Handler: logRouter,
		Addr:    ":8080",
	}

	log.Println("Ready to serve rblog requests ...")
	log.Fatal(server.ListenAndServe())
}
