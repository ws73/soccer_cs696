package main

import (
	// Standard library packages
	"fmt"
	"net/http"

	// Project based packages
	"github.com/ws73@njit.edu/soccer/controllers"

	// Third party library packages
	"github.com/julienschmidt/httprouter"
)

func main() {

	// set the http port
        port := 3000
        host := fmt.Sprintf("localhost:%d", port)

	fmt.Printf("\nRunning at port: %d\n\n", port)	

        // Init the router
        router := httprouter.New()

        // Serve Static Files
        router.ServeFiles("/vendors/*filepath", http.Dir("public/vendors"))
        router.ServeFiles("/resources/*filepath", http.Dir("public/resources"))

        // Setup the app controllers
        controllers.Setup(router)

        // Start the server
        http.ListenAndServe(host, router)
}

