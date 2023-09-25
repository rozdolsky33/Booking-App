package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//// RenderTemplate renders templates using html/template
//func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
//	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.page.html")
//	err := parsedTemplate.Execute(w, nil)
//	if err != nil {
//		fmt.Println("error parsing template:", err)
//		return
//	}
//}

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// check to see if we have the template in out cache
	_, inMap := tc[t]
	if !inMap {
		// need to create the template
		log.Println("creating template and adding to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println("using cached template")
		}
	} else {
		// we have the template in the cache
		log.Println("using cached template")
	}

	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println("using cached template")
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.page.html",
	}
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	// add template to cache (map)
	tc[t] = tmpl
	return nil
}
