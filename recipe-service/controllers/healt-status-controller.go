package controllers

import (
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/core"
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/util"
	"net/http"
)

/*
 * Encapsulating struct to implement controller interface
 */
type HealthStatusController struct {
}

/*
 * This is a simple implementation of the type Controller which will return the request handler function
 * This can be used to check health status while deployment to ascertain the status of app
 */
func (cont *HealthStatusController) Apply(service core.RecipeService) core.RequestContext {
	return func(writer http.ResponseWriter, request *http.Request) {
		util.ResponseWithJSON(writer, http.StatusOK, struct {
			Status string
		}{"ok"})
	}
}
