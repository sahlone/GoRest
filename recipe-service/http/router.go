package http

import (
	"github.com/gorilla/mux"
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/controllers"
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/core"
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/util"
	"net/http"
)

type Router struct {
	*mux.Router
}

/*
 * Define the router routes. Routes will be defined on the Router object rather than on the App object
 *  so when the routes increase we can do the multiple router composition and the app object can have
 * single router as an entry point
 */
func (router *Router) InitializeRoutes(service core.RecipeService) {

	/*
		Authentication controller
	*/
	router.HandleFunc("/auth",
		BindController(&controllers.AuthController{}, nil)).Methods(http.MethodPost)

	/*
		create recipe
	*/
	router.HandleFunc("/recipes",
		BindController(&controllers.CreateRecipeController{}, service)).Methods(http.MethodPost)

	/*
				get recipe list
		 		GET /recipes/{start:[0-9]+}/{limit:[0-9]+} | non-protected
	*/
	router.HandleFunc("/recipes",
		BindController(&controllers.ListRecipeController{}, service)).Methods(http.MethodGet)
	/*
			get single recipe
		 	GET /recipes/{id} | non-protected
	*/
	router.HandleFunc("/recipes/{id}",
		BindController(&controllers.GetRecipeController{}, service)).Methods(http.MethodGet)
	/*
			update recipe
		 	PUT /recipes/{id} | basic auth
	*/
	router.HandleFunc("/recipes/{id}",
		BindController(&controllers.UpdateRecipeController{}, service)).Methods(http.MethodPut)
	/*
			patch recipe
		 	PUT /recipes/{id} | basic auth
	*/
	router.HandleFunc("/recipes/{id}",
		BindController(&controllers.PatchRecipeController{}, service)).Methods(http.MethodPatch)
	/*
			delete recipe
		 	DELETE /recipes/{id} | basic auth
	*/
	router.HandleFunc("/recipes/{id}",
		BindController(&controllers.DeleteRecipeController{}, service)).Methods(http.MethodDelete)
	/*
			rate recipe
		 	PUT /recipes/{id}/rating/{rate:[1-5]} | basic auth
	*/
	router.HandleFunc("/recipes/{id}/rating/{rating}",
		BindController(&controllers.RateRecipeController{}, service)).Methods(http.MethodPut)
	/*
			search recipe by name
		 	GET /recipes/search/{name} | non-protected
	*/
	router.HandleFunc("/recipes/search",
		BindController(&controllers.SearchRecipeController{}, service)).Methods(http.MethodPost)

	/*
		Health Status controller
	*/
	router.HandleFunc("/health",
		BindController(&controllers.HealthStatusController{}, service)).Methods(http.MethodPost, http.MethodGet)
	/*
		Index controller
	*/
	router.HandleFunc("/", util.Index()).Methods(http.MethodPost, http.MethodGet)

}
