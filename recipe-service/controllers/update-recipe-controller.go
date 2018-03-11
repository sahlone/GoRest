package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/core"
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/model"
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/util"
	"net/http"
	"strings"
	"github.com/go-errors/errors"
)

/*
 * Encapsulating struct to implement controller interface
 */
type UpdateRecipeController struct {
}

/*
 * This is a simple implementation of the type Controller which will return the request handler function
 */
func (cont *UpdateRecipeController) Apply(service core.RecipeService) core.RequestContext {
	return func(writer http.ResponseWriter, request *http.Request) {

		if err := util.AuthenticateRequest(writer, request); err != nil {
			return
		}

		vars := mux.Vars(request)
		id := (vars["id"])
		var recipe model.Recipe
		decoder := json.NewDecoder(request.Body)
		defer request.Body.Close()
		err := decoder.Decode(&recipe)
		if err == nil {
			ers:=core.ValidateStruct(recipe)
			if len(ers) >0  {
				err=errors.New(strings.Join(ers,","))
			}
		}
		if  err != nil {
			resp := util.BadRequest(err)
			resp(writer, request)
			return
		}
		recipe.ID = id
		if err := service.UpdateRecipe(recipe); err != nil {
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
