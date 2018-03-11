package controllers

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/core"
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/model"
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/util"
	"net/http"
	"time"
	"errors"
	"strings"
)

/*
 * Encapsulating struct to implement controller interface
 */
type CreateRecipeController struct {
}

/*
 * This is a simple implementation of the type Controller which will return the request handler function
 */
func (cont *CreateRecipeController) Apply(service core.RecipeService) core.RequestContext {
	return func(writer http.ResponseWriter, request *http.Request) {

		if err := util.AuthenticateRequest(writer, request); err != nil {
			return
		}
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
		if err != nil {
			resp := util.BadRequest(err)
			resp(writer, request)
			return
		}
		recipe.ID = uuid.New().String()
		recipe.Modified = time.Now()
		if err := service.CreateRecipe(recipe); err != nil {
			util.ServerError(err)(writer, request)
			return
		}

		util.ResponseWithJSON(writer, http.StatusCreated, recipe)
	}
}
