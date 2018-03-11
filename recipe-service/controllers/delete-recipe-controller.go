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
type DeleteRecipeController struct {
}

/*
 * This is a simple implementation of the type Controller which will return the request handler function
 */
func (cont *DeleteRecipeController) Apply(service core.RecipeService) core.RequestContext {
	return func(writer http.ResponseWriter, request *http.Request) {

		if err := util.AuthenticateRequest(writer, request); err != nil {
			return
		}

		vars := mux.Vars(request)
		id := (core.ID)(vars["id"])
		if err := service.DeleteRecipe(id); err != nil {
			if err == model.ErrNotFound {
				util.NotFound(err)(writer, request)
			} else {
				util.ServerError(err)(writer, request)
			}
			return
		}
		util.ResponseWithJSON(writer, http.StatusNoContent, nil)
	}
}
