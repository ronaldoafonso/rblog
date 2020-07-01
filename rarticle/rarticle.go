/* File: rarticle.go */
/* Description: Handle articles for RBlog. */
package rarticle

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

type article struct {
	lang string
	name string
}

func (a *article) InitTemplates() []string {
	return []string{
		"rarticle/html/" + a.lang + "/layout.html",
		"rarticle/html/" + a.lang + "/head.html",
		"rarticle/html/" + a.lang + "/header.html",
		"rarticle/html/" + a.lang + "/" + a.name + ".html",
		"rarticle/html/" + a.lang + "/footer.html",
	}
}

func (a article) Handler(w http.ResponseWriter, r *http.Request) {
	templates := a.InitTemplates()
	tmplt, err := template.ParseFiles(templates...)
	if err == nil {
		tmplt.ExecuteTemplate(w, "layout", "")
	} else {
		fmt.Println(err)
	}
}

// HandleArticle ... Handle an rblog article.
func HandleArticle(w http.ResponseWriter, r *http.Request) {
	a := article{}
	vars := mux.Vars(r)
	if len(vars) == 2 {
		a.lang = vars["lang"]
		a.name = vars["name"]
	} else if len(vars) == 0 {
		a.lang = "en"
		a.name = "index"
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Page Not Found.")
		return
	}

	a.Handler(w, r)
}
