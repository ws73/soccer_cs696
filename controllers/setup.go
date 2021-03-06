package controllers

import (
	// Standard library packages
	"fmt"
	"net/http"
	"database/sql"
	"strconv"

	// Project based packages
	"github.com/ws73@njit.edu/soccer/models"

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

		db, err := sql.Open("mysql", "root:support$01@/soccer")

		data := make(map[string]interface{})

		if err != nil {
			fmt.Printf("\nError found connecting to the database: %s", err.Error())
		} else {
		
			/*matchModel := models.NewMatchModel(db)
			results := matchModel.GetAllMatches()
			fmt.Printf("\nMatch Records: %v", results)*/
			

			//penaltiesModel := models.NewPenaltiesModel(db)
			//pResults := penaltiesModel.GetAllPenalties()

                        
			//coachModel := models.NewCoachModel(db)
			//coachesResults := coachModel.GetAllCoaches()
			//fmt.Printf("\nCoach Records: %v", coachesResults)
							
			matchSchModel := models.NewMatchSchedulesModel(db)
			matchSchResults := matchSchModel.GetAllMatchSchedules()
			//fmt.Printf("\nMatch Schedules Records: %v", matchSchResults)


			scoreModel := models.NewScoreModel(db)
			scoreResults := scoreModel.GetAllScores()

			data["scoreResults"] = scoreResults
			data["matchSchResults"] = matchSchResults
		}

		 err = renderTemplate("site.layout", "index.html", w, data) 

                if err != nil {
                        fmt.Printf("Error handling template: %s", err)
                }

	})

	//  list all of the rosters
	router.GET("/rosters", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		db, err := sql.Open("mysql", "root:support$01@/soccer")

		data := make(map[string]interface{})

		if err != nil {
			fmt.Printf("\nError found connecting to the database: %s", err.Error())
		} else {

			rostersModel := models.NewRostersModel(db)
                        rostersResults := rostersModel.GetAllRosters()

			matchSchModel := models.NewMatchSchedulesModel(db)
			matchSchResults := matchSchModel.GetAllMatchSchedules()

			data["rostersResults"] = rostersResults
			data["matchSchResults"] = matchSchResults

			err = renderTemplate("site.layout", "rosters.html", w, data)

			if err != nil {
				fmt.Printf("Error handling template: %s", err)
			}      
		}
		
	})

	//  Show player info
	router.GET("/players/:id", func(w http.ResponseWriter, r *http.Request, params httprouter.Params){

		db, err := sql.Open("mysql", "root:support$01@/soccer")

                data := make(map[string]interface{})

                if err != nil {
                        fmt.Printf("\nError found connecting to the database: %s", err.Error())
				
			data["err"] = "Unable to locate player"
                } else {

			idStr := params.ByName("id")
			id, err := strconv.ParseInt(idStr, 10, 64)


			if err == nil {
				playersModel := models.NewPlayerModel(db)
				playerInfo, err := playersModel.GetPlayer(id)
				data["playerInfo"] = playerInfo

				if err != nil {
					data["err"] = "Unable to locate player - invalid info may have been provided"
				}
			}else{

				if err != nil {
					data["err"] = "Unable to locate player - invalid info may have been provided"
				}

			}
			

		}	

		matchSchModel := models.NewMatchSchedulesModel(db)
		matchSchResults := matchSchModel.GetAllMatchSchedules()

		data["matchSchResults"] = matchSchResults

		err = renderTemplate("site.layout", "player.html", w, data)

		if err != nil {
			fmt.Printf("Error handling template: %s", err)
		}    

	})

}

