package render

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.html")
	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		log.Println(err)
		return
	}
}

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {

	var tmpl *template.Template
	var err error

	// if we allready have the template in our cache
	_, inMap := tc[t]
	if !inMap {
		log.Println("Created Cached Tempaltes", t)
		err = createTemplateCached(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		log.Println("Using Cached Tempaltes", t)
	}

	tmpl = tc[t]

	err = tmpl.Execute(w, nil)

}

func createTemplateCached(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.html",
	}

	tmpl, err := template.ParseFiles(templates...)

	if err != nil {
		return err
	}

	tc[t] = tmpl

	return nil
}
