package util

import (
	"encoding/json"
	"github.com/go-errors/errors"
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/auth"
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/logger"
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/model"
	"net/http"
	"strings"
)

// ResponseWithJSON response with json
func ResponseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		logger.Info("Error json encode :", err)
	}
	logger.Info("Response sent successfully %+v", payload)
	w.Header().Set("Content-Type", "application/json")

}

func Status204() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ResponseWithJSON(w, http.StatusNoContent, struct {
		}{})
	}
}

func NotFound(v ...interface{}) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ResponseWithJSON(w, http.StatusNotFound, v)
	}
}

func BadRequest(v ...interface{}) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ResponseWithJSON(w, http.StatusBadRequest, v)
	}
}

func Unauthorized(v ...interface{}) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ResponseWithJSON(w, http.StatusUnauthorized, v)
	}
}
func Index() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ResponseWithJSON(w, http.StatusOK, struct {
		}{})
	}
}
func ServerError(err error) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Error("Error serving request :%v", err)
		ResponseWithJSON(w, http.StatusInternalServerError, model.ErrorServiceUnavailable)
	}
}

var err = errors.New("auth.failed")

func AuthenticateRequest(w http.ResponseWriter, r *http.Request) error {

	token := r.Header.Get("Authorization")
	if token == "" {
		Unauthorized()(w, r)
		return err
	}
	values := strings.Split(token, " ")
	if len(values) != 2 {
		Unauthorized()(w, r)
		return err
	}
	if !strings.EqualFold("Bearer", values[0]) {
		Unauthorized()(w, r)
		return err
	}
	err := auth.ValidateToken(values[1])
	if err != nil {
		Unauthorized(err)(w, r)
		return err
	}
	return nil
}
