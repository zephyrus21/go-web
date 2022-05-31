package handlers

import (
	"basic-web/pkg/renders"
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "About page!")
}
