package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html")
}

func Features(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "features.html")
}

func Faqs(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "faqs.html")
}

func Login(w http.ResponseWriter, r *http.Request) {

	// POST isteği ise parametreleri al ?username=xxx&password=xxx
	// if r.Method == "POST" {
	// 	r.ParseForm()
	// 	username := r.FormValue("username")
	// 	password := r.FormValue("password")
	// 	if username == "admin" && password == "admin" {
	// 		http.Redirect(w, r, "/home", http.StatusFound)
	// 	} else {
	// 		http.Error(w, "Geçersiz kimlik bilgileri", http.StatusUnauthorized)
	// 		http.Redirect(w, r, "/login", http.StatusFound)
	// 	}
	// }
	type Credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}

		// JSON verisini Credentials yapısına dönüştür

		var creds Credentials
		err = json.Unmarshal(body, &creds)
		if err != nil {
			http.Error(w, "Error decoding JSON", http.StatusBadRequest)
			return
		}

		// Kullanıcı adı ve şifreyi kontrol et

		if creds.Username == "test" && creds.Password == "test" {
			log.Println("Login Succes {username: " + creds.Username + ", password: " + creds.Password + ", status:" + http.StatusText(http.StatusOK) + "}")
			w.Write([]byte("Giriş başarılı!"))
		} else {
			log.Println("Login Error {username: " + creds.Username + ", password: " + creds.Password + ", status:" + http.StatusText(http.StatusUnauthorized) + "}")
			http.Error(w, "Geçersiz kimlik bilgileri", http.StatusUnauthorized)
		}

		if r.Method == "GET" {
			renderTemplate(w, "login.html")
		}

	}
}
