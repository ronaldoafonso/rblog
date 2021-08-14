/* File: rarticle.go */
/* Description: Handle articles for RBlog. */
package rarticle

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"os"
)

var (
	host = "localhost"
)

type article struct {
	Lang      string
	Name      string
	Host      string
	templates []string
}

func (a *article) InitTemplates() {
	a.templates = []string{
		"rarticle/html/" + a.Lang + "/layout.html",
		"rarticle/html/" + a.Lang + "/head.html",
		"rarticle/html/" + a.Lang + "/header.html",
		"rarticle/html/" + a.Lang + "/" + a.Name + ".html",
		"rarticle/html/" + a.Lang + "/subscribe.html",
		"rarticle/html/" + a.Lang + "/footer.html",
	}
}

func (a *article) ExecuteTemplate(w http.ResponseWriter) {
	tpl, err := template.ParseFiles(a.templates...)
	if err != nil {
		fmt.Println(err)
	} else {
		tpl.ExecuteTemplate(w, "layout", a)
	}
}

func (a article) Handler(w http.ResponseWriter, r *http.Request) {
	a.InitTemplates()
	a.ExecuteTemplate(w)
}

// HandleArticle ... Handle an rblog article.
func HandleArticle(w http.ResponseWriter, r *http.Request) {
	a := article{Host: host}
	vars := mux.Vars(r)
	if len(vars) == 2 {
		a.Lang = vars["lang"]
		a.Name = vars["name"]
	} else if len(vars) == 0 {
		a.Lang = "en"
		a.Name = "index"
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Page Not Found.")
		return
	}
	a.Handler(w, r)
}

func HandleSubscriber(w http.ResponseWriter, r *http.Request) {
	lang := mux.Vars(r)["lang"]
	subscriber := r.FormValue("subscriber")
	email := r.FormValue("email")

	f, err := os.OpenFile("/tmp/subscriber.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	defer f.Close()

	fmt.Fprintf(f, "%v,%v,%v\n", subscriber, email, lang)

	a := article{
		Name: "subscribed",
		Lang: lang,
		Host: host,
	}
	a.Handler(w, r)
}

// InitHost ... Init variable host for the rarticle package.
func InitHost(h *string) {
	host = *h
}
