package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// "index" page for rblog.
func index(w http.ResponseWriter, r *http.Request) {
	indexContent := []string{
		"Hi there!",
		"My name is Ronaldo Afonso.",
		"I've been working in the Linux and network fields for the last 20 years.",
		"I first met Linux when I installed an Slackware distro version 3.4.",
		"I can't even remember exactly when it was, but I would say that it was in 1999.",
		"At that time, there weren't plenty of documentation as there are today.",
		"One day a friend of mine just lent me a Linux Magazine with a bundled Slackware CD-ROM.",
		"So I tried to install it on my computer and guess what? After three months of trying, I did it.",
		"I've been using Linux since then.",
		"This blog is about my experience with Linux, network and programming.",
		"I really hope you enjoy it.",
		"[]s",
		"Ronaldo Afonso",
	}

	if indexTemplate, err := template.ParseFiles("index.html"); err == nil {
		indexTemplate.Execute(w, indexContent)
	} else {
		fmt.Println(err)
	}
}

func main() {
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}

	http.HandleFunc("/", index)

	fmt.Println("Ready to serve rblog requests ...")
	server.ListenAndServe()
}
