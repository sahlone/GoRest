package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	. "github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/core"
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/model"
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/util"
	"net/http"
	"strings"
	"github.com/go-errors/errors"
)

/*
 * Encapsulating struct to implement controller interface
 */
type PatchRecipeController struct {
}

/*
 * This is a simple implementation of the type Controller which will return the request handler function
 */
func (cont *PatchRecipeController) Apply(service RecipeService) RequestContext {
	return func(writer http.ResponseWriter, request *http.Request) {

		if err := util.AuthenticateRequest(writer, request); err != nil {
			return
		}

		vars := mux.Vars(request)
		id := (ID)(vars["id"])
		var recipe model.RecipePatch
		decoder := json.NewDecoder(request.Body)
		defer request.Body.Close()
		err := decoder.Decode(&recipe)
		if err == nil {
			ers:= ValidateStruct(recipe)
			if len(ers) >0  {
				err=errors.New(strings.Join(ers,","))
			}
		}
		if  err != nil {
			resp := util.BadRequest(model.ErrorInvalidRequest)
			resp(writer, request)
			return
		}

		if err := service.PatchRecipe(id, recipe); err != nil {
			switch err.(type) {
			case model.Error:
				if(err==model.ErrNotFound){
					util.NotFound(err)(writer, request)
				}else{
					util.BadRequest(err)(writer, request)
				}
			default:
				util.ServerError(err)(writer, request)
			}

			return
		}

		util.ResponseWithJSON(writer, http.StatusNoContent, nil)
	}
}
