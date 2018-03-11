package http

import (
	"fmt"
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/core"
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/logger"
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/util"
	"net/http"
	"reflect"
)

/*
 * The Binder has been defined as an intermediate step to use the abstraction provided by the Controller interface
 * It will also act as a decorator design pattern where we do pre app initialization steps and pre request steps like logging /filter per
 * request execution
 * Like in the code below we have handled the global 500. This will also prevent any panics if thrown from below layers
 */

func BindController(cont core.Controller, service core.RecipeService) func(w http.ResponseWriter, r *http.Request) {
	logger.Info("Binding controller %v ", reflect.TypeOf(cont))
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Request Received :%+v", *r)
		defer func() {
			if err := recover(); err != nil {
				logger.Info("Error serving request:")
				logger.Info(fmt.Sprint(err))
				util.ServerError(err.(error))(w, r)
			}
		}()
		cont.Apply(service)(w, r)

	}
}
