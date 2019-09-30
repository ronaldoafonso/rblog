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
	indexContent := []string{
		"Hi there!",
		"My name is Ronaldo Afonso.",
		"I've been working in the Linux and network fields for the last 20 years.",
		"I first met Linux when I installed an Slackware distro version 3.4.",
		"I can't even remember exactly when it was, but I would say that it was in 1999.",
		"At that time, there wasn't plenty of documentation as there is today.",
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

// "about" page for rblog.
func about(w http.ResponseWriter, r *http.Request) {
	aboutContent := []string{
		"I started my carrer in 1998.",
		"I was a 16-year boy looking for books about computers and ended up finding one about C programming.",
		"That book really got my atention, so I decided to give it a try.",
		"Well, I had no idea how that C programming stuff would help me, but I thought that C programmers know a lot about computers.",
		"After five months (yes that's it, I spent almost half a year to get a C compiler up and running), I was able to run my first \"Hello World!\" in C.",
		"I didn't have access to the Internet like we do today. I had to dail to my Internet Service Provider in order to get a PPP connection.",
		"Some times I asked friends or people from the computer industry from my home town for help or to lend me some CD-ROMs.",
		"That was how I got computer related contents.",
		"One year went by until I decided I would like to be good at Linux and network.",
		"Once more I asked a friend to help me. That friend lend me a CD-ROM with a bundled Slackware CD-ROM.",
		"Three months later I got my first Linux installation and could start to learn C programming using the Linux operating system.",
		"My dream at that time was to get a job as a C programming. I got one, but my boss told me that I needed to install a Linux server with an Apache and MySQL server first.",
		"So I did it, but when I was just about to be \"promoted\", I quit that job and started another one as a kind of Linux network consulting.",
		"Yes, it was a big step for me.",
		"I had just installed my first Linux distrubution one year latter and had become a Linux Network Consultant.",
		"Well, I think I did good, but I must confess, I did a lot of things wrongs. But it was good for my education in the computer field.",
		"After three years at the position, I finnally started a job as a C programmer. I was one of the main developer for a suite of router that run Linux as their operationg system.",
		"Working in that role I learned a lot about Linux, network and of couse, C programming.",
		"After that I quited and started a new job as C programming too, but this time I had to be responsible for almost 1000 Wi-Fi routers.",
		"And that what I'm doing until now.",
		"Today I'm in charge of more than 2000 Wi-Fi router that run Linux (OpenWrt).",
	}

	if aboutTemplate, err := template.ParseFiles("about.html"); err == nil {
		aboutTemplate.Execute(w, aboutContent)
	} else {
		fmt.Println(err)
	}
}
