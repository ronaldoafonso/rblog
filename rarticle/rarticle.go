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
	templates []string = []string{
		"rarticle/html/layout.html",
		"rarticle/html/head.html",
		"rarticle/html/header.html",
		"rarticle/html/index.html",
		"rarticle/html/footer.html",
	}
	articles map[string]func(http.ResponseWriter, *http.Request)
)

func init() {
	articles = make(map[string]func(http.ResponseWriter, *http.Request))
	articles["index"] = index
	articles["about"] = about
}

// HandleArticle ... Handle an rblog article.
func HandleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if len(vars) == 0 {
		articles["index"](w, r)
	} else if article, ok := articles[vars["article"]]; ok {
		article(w, r)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Page Not Found.")
	}
}

// index ... Handle index page.
func index(w http.ResponseWriter, r *http.Request) {
	templates[3] = "rarticle/html/index.html"
	indexTemplate, err := template.ParseFiles(templates...)
	if err == nil {
		indexTemplate.ExecuteTemplate(w, "layout", "")
	} else {
		fmt.Println(err)
	}
}

// about .. Handle "about" page.
func about(w http.ResponseWriter, r *http.Request) {
	templates[3] = "rarticle/html/about.html"
	aboutTemplate, err := template.ParseFiles(templates...)
	if err == nil {
		aboutTemplate.ExecuteTemplate(w, "layout", "")
	} else {
		fmt.Println(err)
	}
}
