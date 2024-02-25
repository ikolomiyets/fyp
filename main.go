package main

import (
	"FYP/db"
	"FYP/handlers"
	"FYP/oauth2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"net/http"
	"os"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "*",
		AllowHeaders: "*",
	}))

	dbClient := db.MustCreate(os.Getenv("DB_URL"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD")) // create db client
	controller := handlers.New(dbClient)                                                               //dependancy injection

	// Initialize oauth2 middleware
	oauth2Config, err := oauth2.Build(
		oauth2.Debug(true),
		oauth2.URL(os.Getenv("JWKS_URL")),
		oauth2.Unmatched(true),
		oauth2.Audience(os.Getenv("AUDIENCE")),
		oauth2.Issuer(os.Getenv("ISSUER")),
		oauth2.HTTPClient(&http.Client{}),
		oauth2.Request("GET", "/questions", []string{"read:questions"}),
	)
	if err != nil {
		log.Println("cannot create OAuth2 middleware")
		os.Exit(2)
	}

	app.Use(oauth2.New(oauth2Config))

	app.Post("/authorize", controller.AuthorizeHandler)
	app.Post("/newQuestion", controller.NewQuestion) //creates new question
	app.Post("/newAnswer", controller.NewAnswer)     //creates new answer for particular question and adds to db
	app.Get("/questions", controller.GetQuestionsHandler)
	//app.Get("/isSupervisor", controller.GetSupervisorHandler)
	app.Get("getApplications", controller.GetApplicationsHandler)                     //retrieves all applications from db
	app.Get("getSpecificApplications/:id", controller.GetSpecificApplicationsHandler) //retrieves one specific applications
	app.Get("/getGanttItem/:id", controller.GetGanttItem)
	app.Get("/getGantt/:id", controller.GetGantt)
	app.Get("/getSupervisors", controller.GetSupervisorHandler)
	//getapplications
	//getspecificapplication
	//post newquestion
	//post newanswer
	//post createproject
	//post createapplication
	//post acceptapplication
	//post declineapplication
	//post createganttitem
	//post addfeedback
	app.Listen(":3000")
}
