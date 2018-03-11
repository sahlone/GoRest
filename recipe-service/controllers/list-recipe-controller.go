package controllers

import (
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/core"
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/model"
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/util"

	"net/http"
	"strconv"
)

/*
 * Encapsulating struct to implement controller interface
 */
type ListRecipeController struct {
}

/*
 * This is a simple implementation of the type Controller which will return the request handler function
 */
func (cont *ListRecipeController) Apply(service core.RecipeService) core.RequestContext {
	return func(writer http.ResponseWriter, request *http.Request) {
		limit := request.URL.Query().Get("limit")
		lastRecord := request.URL.Query().Get("lastRecord")

		var limitInt int64
		var err error
		if len(limit) > 0 {
			limitInt, err = strconv.ParseInt(limit, 10, 32)
			if err != nil {
				util.BadRequest(model.ErrorGetLimit)(writer, request)
				return
			}
		} else {
			limitInt = 10
		}

		recipes, err := service.GetRecipes(lastRecord, int(limitInt))
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
