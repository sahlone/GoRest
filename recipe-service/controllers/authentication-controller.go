package controllers

import (
	"encoding/json"
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/auth"
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/core"
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/model"
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/util"
	"net/http"
)

/*
 * Encapsulating struct to implement controller interface
 */
type AuthController struct {
}

/*
 * This is a simple implementation of the type Controller which will return the request handler function
 */
func (cont *AuthController) Apply(service core.RecipeService) core.RequestContext {
	return func(writer http.ResponseWriter, request *http.Request) {

		creds := model.Creds{}
		decoder := json.NewDecoder(request.Body)
		defer request.Body.Close()
		if err := decoder.Decode(&creds); err != nil {
			resp := util.BadRequest(model.ErrCredsInvalid)
			resp(writer, request)
			return
		}
		jwtToken, err := auth.Authorize(creds.User, creds.Password)
		if err != nil {
			util.Unauthorized(err)(writer, request)
			return
		}

		util.ResponseWithJSON(writer, http.StatusOK, model.IDToken{jwtToken})
	}
}
