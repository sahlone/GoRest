package http

import (
	"github.com/gorilla/mux"
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/config"
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/core"
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/logger"
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/modules"
	"net/http"
	"time"
)

/*
 * Application container
 */
type App struct {
	*Router
	core.RecipeService
	config.Config
}

/*
 * Represents the formation of App struct which is used to wire the components of the App
 */
func Initialize() App {
	config := config.GetConfig()
	handler := modules.NewMongoHandler(&config)
	recipeRepo := modules.NewRecipeDBRepo(handler)
	service := modules.NewRecipeService(recipeRepo)

	app := App{&Router{mux.NewRouter()}, service, config}
	app.InitializeRoutes(service)
	return app
}

/*
 * Entry point for the App struct to start listening on a address
 * Multiple App containers can be used to start multiple servers service different routes
 * The startup is made fail fast so we dont start any container if there is any problem
 */
func (app *App) Run(addr string) {
	srv := &http.Server{
		Handler:      app.Router,
		Addr:         addr,
		WriteTimeout: time.Duration(app.WriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(app.ReadTimeout) * time.Second,
	}
	err := srv.ListenAndServe()
	if err != nil {
		logger.Error("Error in starting the app container %v", err)
		logger.Fatal(err)
	}

}
