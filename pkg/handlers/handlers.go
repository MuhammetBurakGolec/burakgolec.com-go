package handlers

import (
	"net/http"

	"github.com/muhammetburakgolec/burakgolec.com-go/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.html")
}

func Features(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "features.html")
}

func Faqs(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "faqs.html")
}
