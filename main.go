package main

import (
	"fmt"
	"net/http"
)

// "index" page for rblog.
func index(w http.ResponseWriter, r *http.Request) {
	htmlIndex := `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <meta name="google-site-verification"
        content="oSH_M4dBKpTHP3hOAmXeRbc6RfwNqovYL9mBKqgKQJs" />
  <link rel="stylesheet"
        href="https://cdn.ronaldoafonso.com.br/css/ra-main.css">
  <title>Ronaldo Afonso</title>
</head>
<body>
  <header>
    <img src="https://cdn.ronaldoafonso.com.br/ra-logo.png"
         alt="Ronaldo Afonso's Logo">
    <h1>Ronaldo Afonso - Linux, network and some programming</h1>
  </header>
  <main>
    <p>Hi!</p>
    <p>My name is Ronaldo Afonso.</p>
    <p>I've been working in the Linux and network fields for the last 20 years.</p>
    <p>I first met Linux when I installed an Slackware 3.4 distribution.</p>
    <p>I can't exactly remember when it was, but I would say that it was in 1999.</p>
    <p>There wasn't a plenty of documentation that time as there is today.</p>
    <p>A friend of mine just lent me a Linux Magazine with a Slackware CD-ROM. I tried to install it on my computer and guess what? I did it.</p>
    <p>I've been using Linux since then.</p>
    <p>This blog is about my experiences with Linux, network and programming.</p>
    <p>I really hope you enjoy it.<p>
    <p>[]s</p>
    <p>Ronaldo Afonso</p>
  </main>
  <footer>
    <small>
    <a href="mailto:ronaldo@ronaldoafonso.com.br">Email me</a> |
    <a href="https://www.linkedin.com/in/ronaldoafonso/" target="_blank">LinkedIn</a> |
    <a href="https://github.com/ronaldoafonso" target="_blank">GitHub</a> |
    <a href="https://twitter.com/ronaldoafonso" target="_blank">Twitter</a>
    </small>
  </footer>
</body>
</html>`
	fmt.Fprintf(w, htmlIndex)
}

func main() {
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}

	http.HandleFunc("/", index)

	fmt.Println("Ready to serve rblog requests ...")
	server.ListenAndServe()
}
