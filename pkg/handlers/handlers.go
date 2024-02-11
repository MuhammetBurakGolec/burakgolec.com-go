package handlers

import (
	"log"
	"net/http"

	"github.com/muhammetburakgolec/burakgolec.com-go/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	ip_track(r)
	render.RenderTemplate(w, "home.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	ip_track(r)
	render.RenderTemplate(w, "about.html")
}

func Features(w http.ResponseWriter, r *http.Request) {
	ip_track(r)
	render.RenderTemplate(w, "features.html")
}

func Faqs(w http.ResponseWriter, r *http.Request) {
	ip_track(r)
	render.RenderTemplate(w, "faqs.html")
}

func ip_track(r *http.Request) {
	ip := r.RemoteAddr
	if forwarded := r.Header.Get("X-Forwarded-For"); forwarded != "" {
		ip = forwarded
	}
	log.Printf("Gelen istek IP: %s\n", ip)
}
