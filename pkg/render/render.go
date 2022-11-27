package render

import (
	"bytes"
	"log"
	"github.com/Ramnath-Karthikesan/bookings/pkg/config"
	"github.com/Ramnath-Karthikesan/bookings/pkg/models"
	"net/http"
	"path/filepath"
	"text/template"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig){
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData{
	return td
}

// RenderTemplate renders template using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	var tc map[string]*template.Template
	// get the template cache from the app config
	if app.UseCache {
		tc = app.TemplateCache
	}else {
		tc, _ = CreateTemplateCache()
	}
	

	// create a template cache
	// tc, err := CreateTemplateCache()

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// get the requested template

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}
	td = AddDefaultData(td)

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, td)
	// if err != nil {
	// 	log.Println(err)
	// }

	// render the template

	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
	// parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.html")
	// err := parsedTemplate.Execute(w, nil)
	// if err != nil {
	// 	fmt.Println("error parsing template:", err)
	// 	return
	// }
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all files ending with .page.tmpl from ./templates

	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")

		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err := ts.ParseGlob("./templates/*.layout.tmpl")

			if err != nil {
				return myCache, err
			}

			myCache[name] = ts
		}

		myCache[name] = ts
		// log.Println(myCache)
	}

	return myCache, nil
}

// var tc = make(map[string]*template.Template)

// func RenderTemplate(w http.ResponseWriter, t string) {
// 	var tmpl *template.Template
// 	var err error

// 	//check to see if we already have the template in our cache
// 	_, inMap := tc[t]
// 	log.Println(t)
// 	log.Println(inMap)
// 	log.Println(1)
// 	if !inMap {
// 		//need to create the template
// 		log.Println(2)
// 		log.Println("creating template and adding to cache")
// 		err = createTemplateCache(t)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	} else {
// 		//we have the template in the cache
// 		log.Println(3)
// 		log.Println("using cached template")
// 	}
// 	log.Println(4)
// 	tmpl = tc[t]
// 	err = tmpl.Execute(w, nil)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// func createTemplateCache(t string) error {
// 	// templates := []string{
// 	// 	fmt.Sprintf("./templates/%s", t),
// 	// 	"./templates/base.html",
// 	// }

// 	//pare the template
// 	tmpl, err := template.ParseFiles("./templates/"+t, "./templates/base.html")
// 	if err != nil {
// 		return err
// 	}

// 	tc[t] = tmpl
// 	return nil
// }
