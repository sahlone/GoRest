package controllers

import (
	"encoding/json"
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/core"
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/model"
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/util"
	"net/http"
)

/*
 * Encapsulating struct to implement controller interface
 */
type SearchRecipeController struct {
}

/*
 * This is a simple implementation of the type Controller which will return the request handler function
 */
func (cont *SearchRecipeController) Apply(service core.RecipeService) core.RequestContext {
	return func(writer http.ResponseWriter, request *http.Request) {
		var criteria model.SearchCriteria
		decoder := json.NewDecoder(request.Body)
		defer request.Body.Close()
		if err := decoder.Decode(&criteria); err != nil {
			resp := util.BadRequest(model.ErrorInvalidRequest)
			resp(writer, request)
			return
		}
		recipes, err := service.SearchRecipes(criteria.Start, criteria.Limit, criteria)
		if err != nil {
			switch err.(type) {

			case model.Error:
				util.BadRequest(err)(writer, request)
			default:
				util.ServerError(err)(writer, request)
			}
			return
		}

		util.ResponseWithJSON(writer, http.StatusOK, recipes)
	}
}
