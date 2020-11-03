package main

import (
	"flag"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/ronaldoafonso/rblog/rarticle"
	"log"
	"net/http"
	"os"
)

func main() {
	host := flag.String("host", "localhost", "URL host")
	flag.Parse()

	router := mux.NewRouter()
	router.HandleFunc("/", rarticle.HandleArticle)
	router.HandleFunc("/{lang}/{name}", rarticle.HandleArticle)
	logRouter := handlers.LoggingHandler(os.Stdout, router)

	server := &http.Server{
		Handler: logRouter,
		Addr:    ":8080",
	}

	rarticle.InitHost(host)

	log.Println("Ready to serve rblog requests ...")
	log.Fatal(server.ListenAndServe())
}
