/* File: rarticle.go */
/* Description: Handle articles for RBlog. */
package rarticle

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

var (
	host = "localhost"
)

type article struct {
	Lang string
	Name string
	tpls []string
	Host string
}

func (a *article) InitTemplates() {
	a.tpls = []string{
		"rarticle/html/" + a.Lang + "/layout.html",
		"rarticle/html/" + a.Lang + "/head.html",
		"rarticle/html/" + a.Lang + "/header.html",
		"rarticle/html/" + a.Lang + "/" + a.Name + ".html",
		"rarticle/html/" + a.Lang + "/footer.html",
	}
}

func (a *article) ExecuteTemplate(w http.ResponseWriter) {
	tpl, err := template.ParseFiles(a.tpls...)
	if err == nil {
		tpl.ExecuteTemplate(w, "layout", a)
	} else {
		fmt.Println(err)
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

// InitHost ... Init variable host for the rarticle package.
func InitHost(h *string) {
	host = *h
}
