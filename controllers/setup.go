package controllers

import (
	// Standard library packages
	"fmt"
	"net/http"

	// Third party package libraries
	"github.com/julienschmidt/httprouter"
)

// Setup has the base instantiations for the controllers.
// It sets up the basic and complex routers for the website
func Setup(router *httprouter.Router) {

	// load templates
	loadTemplates()

	// call the router handlers
	setRouterHandlers(router)
}

// setRouterHandlers is a helper function that registers all of routers for the site
func setRouterHandlers(router *httprouter.Router) {

	// ===================================================================
	// Basic Controllers are handled below
	// ===================================================================

	// home page controller
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		err := renderTemplate("site.layout", "index.html", w, "")

		if err != nil {
			fmt.Printf("Error handling template: %s", err)
		}
	})
}
