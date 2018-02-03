package controllers

import (
	// Standard Library Packages
	"fmt"
	"html/template"
	"net/http"
	//"os"
	"path/filepath"
)

var appTemplates map[string]*template.Template

// panicOnError runs a panic when error is not nil
func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

// LoadTemplates loads the app templates or pages that are used throughout the application
func loadTemplates() {

	fmt.Println("About to load templates")

	// get layouts
	layouts, err := filepath.Glob("templates/layouts/*.layout")
	panicOnError(err)

	// get list of main pages
	pages, err := filepath.Glob("templates/pages/*.html")
	panicOnError(err)

	for _, page := range pages {
		files := append(layouts, page)
		templateName := filepath.Base(page)

		newTemplate := template.Must(template.ParseFiles(files...))
		newTemplate.Option("missingkey=default")

		appTemplates[templateName] = newTemplate
	}

	// loaded templates
	for file, _ := range appTemplates {
		fmt.Printf("Loaded Template: %s\n", file)
		fmt.Printf("loaded: %s\n", file)
	}

}

// renderTemplate writes the template to the http.ResponseWriter
func renderTemplate(layout, name string, w http.ResponseWriter, data interface{}) error {

	t, ok := appTemplates[name]

	if !ok {
		err := fmt.Errorf("Unable to find the template: %s\n", name)

		fmt.Println(err.Error())

		return err
	}

	err := t.ExecuteTemplate(w, layout, data)

	return err
}


func init (){
	// holds the templates for the project
	appTemplates = make(map[string]*template.Template)
}
