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
	articles map[string]article = make(map[string]article)
)

type article struct {
    name string
}

func (a article) Handler(w http.ResponseWriter, r *http.Request) {
	templates[3] = "rarticle/html/"+a.name+".html"
	tmplt, err := template.ParseFiles(templates...)
	if err == nil {
		tmplt.ExecuteTemplate(w, "layout", "")
	} else {
		fmt.Println(err)
	}
}

func init() {
	articles["index"] = article{"index"}
	articles["about"] = article{"about"}
}

// HandleArticle ... Handle an rblog article.
func HandleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if len(vars) == 0 {
		articles["index"].Handler(w, r)
	} else if a, ok := articles[vars["article"]]; ok {
		a.Handler(w, r)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Page Not Found.")
	}
}
