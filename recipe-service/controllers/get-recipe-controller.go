package controllers

import (
	"github.com/gorilla/mux"
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/core"
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/model"
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/util"
	"net/http"
)

/*
 * Encapsulating struct to implement controller interface
 */
type GetRecipeController struct {
}

/*
 * This is a simple implementation of the type Controller which will return the request handler function
 */
func (cont *GetRecipeController) Apply(service core.RecipeService) core.RequestContext {
	return func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		id := (core.ID)(vars["id"])
		recipe, err := service.GetRecipe(id)
		if err != nil {
			switch err.(type) {

			case model.Error:
				util.NotFound(err)(writer, request)
			default:
				util.ServerError(err)(writer, request)
			}

			return
		}
		util.ResponseWithJSON(writer, http.StatusOK, recipe)
	}
}
