package handlers

import (
	// "fmt"
	// "html/template"
	"net/http"
	"github.com/sayheyjitendra/learn-go/package/render"
)



func Home (w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, "home.page.tmpl")

}
func About (w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, "about.page.tmpl")
}



