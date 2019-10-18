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

// "about" page for rblog.
func about(w http.ResponseWriter, r *http.Request) {
	aboutContent := []string{
		"I started my computer journey in about 1998.",
		"I was a 16-year-old boy looking for books about computers and ended up finding one about C programming language.",
		"That book really got my attention, so I decided to give it a try.",
		"Well, I had no idea how that C programming language stuff would help me, but somehow I believed that C programmers know a lot about computers.",
		"After five months (yes that's it, I spent almost half a year to get a C compiler up and running), I was able to run my first \"Hello World!\" in C.",
		"I didn't have Internet access like we do today and when I did, it had to be as a dial to an Internet Service Provider. Something known as PPP.",
		"Some times I asked friends or people from the computer industry from my home town for help or to lend me some CD-ROMs, books, magazines, etc.",
		"That was how I could have access to some computer related content at that time.",
		"One year went by and I heard about Linux. It sounded like a very nice operating system. So, I decided I would like to be good at it.",
		"Once more I asked a friend to help me. That friend lend me a CD-ROM with a bundled Slackware 3.4 in it.",
		"Three months later I got my first Linux installation and could start to learn C programming using the Linux operating system.",
		"My dream at that time was to get a job as a C programmer. I got one, but my boss told me that I needed to install a Linux server with an Apache and MySQL server first.",
		"So I did it, but when I was just about to be \"promoted\", I had to quit that job and started another one as a kind of Linux Network Administrator.",
		"Yeah, that was a big step for me. It had just been passed one year since I installed my first Linux distributioni, and even with so little time of experience I was hired as a Linux Network Administrator.",
		"Well, I think I did good. But I must confess, I made a lot of mistakes. In some way, it was good for my education in the computer field.",
		"After three years working as a Linux Network Administrator, I finally started a job as a C programmer.",
		"It was a big step for me. I was hired to be one of the C programmers responsible for a suite of routers which run Linux as their operating system.",
		"I would say that working in the role I had an opportunity to really learn a lot about Linux, computer network and of course, C programming.",
		"After about three years I just quited that job and started working for a Wi-Fi Internet Service Provided.",
		"Again as C programming, but this time I was also responsible for managing a network of Wi-Fi routers.",
		"And that's what I'm doing until now.",
		"Today I'm in charge of more than 1000 Wi-Fi routers which run the OpenWrt Linux system.",
		"[]s",
		"Ronaldo Afonso",
	}

	if aboutTemplate, err := template.ParseFiles("about.html"); err == nil {
		aboutTemplate.Execute(w, aboutContent)
	} else {
		fmt.Println(err)
	}
}
