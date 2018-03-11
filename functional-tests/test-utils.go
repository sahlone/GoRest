package functional_tests

import (
	"encoding/json"
	"fmt"
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/auth"
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/model"
	. "github.com/onsi/ginkgo"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var serviceUrl = "http://localhost:9000"
var currentRecipeID string = ""
var clientTimeout = time.Duration(1000000 * time.Second)

var testRecipe = model.Recipe{
	Name:       "Test",
	Prep:       model.Preparation{1, 1},
	Difficulty: 1,
	IsVeg:      true,
}

func SetAuthentication(req *http.Request) {

	token, err := auth.Authorize("admin", "password")
	if err != nil {
		Fail("Failed to create authorization token" + err.Error())
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	return
}

func CreateRecipe(recipe model.Recipe, setAuth bool) (model.Recipe, *http.Response) {

	result := model.Recipe{}
	params, _ := json.Marshal(recipe)
	paramstr := string(params)
	req, _ := http.NewRequest("POST", fmt.Sprintf("%v/recipes", serviceUrl), strings.NewReader(paramstr))
	if setAuth {
		SetAuthentication(req)
	}
	req.Header.Set("Accept", "application/json")

	client := &http.Client{Timeout: clientTimeout}
	res, err := client.Do(req)

	if err != nil {
		Fail(err.Error())
	} else {
		defer res.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(res.Body)
		json.Unmarshal(bodyBytes, &result)
	}
	return result, res

}

func GetRecipe(id string) (model.Recipe, *http.Response) {

	result := model.Recipe{}

	req, _ := http.NewRequest("GET", fmt.Sprintf("%v/recipes/%v", serviceUrl, id), nil)
	req.Header.Set("Accept", "application/json")
	client := &http.Client{Timeout: clientTimeout}
	res, err := client.Do(req)

	if err != nil {
		Fail(err.Error())
	} else {
		defer res.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(res.Body)
		json.Unmarshal(bodyBytes, &result)
	}
	return result, res

}

func DeleteRecipe(id string, setAuth bool) *http.Response {

	req, _ := http.NewRequest("DELETE", fmt.Sprintf("%v/recipes/%v", serviceUrl, id), nil)
	req.Header.Set("Accept", "application/json")
	client := &http.Client{Timeout: clientTimeout}
	if setAuth {
		SetAuthentication(req)
	}
	res, err := client.Do(req)
	if err != nil {
		Fail(err.Error())
	}
	return res

}

func GetIDToken(user, pass string) (string, *http.Response) {

	creds := model.Creds{user, pass}

	params, _ := json.Marshal(creds)
	paramstr := string(params)
	req, _ := http.NewRequest("POST", fmt.Sprintf("%v/auth", serviceUrl), strings.NewReader(paramstr))
	req.Header.Set("Accept", "application/json")
	client := &http.Client{Timeout: clientTimeout}
	res, err := client.Do(req)

	token := model.IDToken{}
	if err != nil {
		Fail(err.Error())
	} else {
		defer res.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(res.Body)
		json.Unmarshal(bodyBytes, &token)
	}
	return token.Token, res

}
func ExecuteRequest(method, path, payload string, setAuth bool) *http.Response {

	req, _ := http.NewRequest(method, fmt.Sprintf("%v/%v", serviceUrl, path), strings.NewReader(payload))
	req.Header.Set("Accept", "application/json")
	client := &http.Client{Timeout: clientTimeout}
	if setAuth {
		SetAuthentication(req)
	}
	res, err := client.Do(req)

	if err != nil {
		Fail(err.Error())
	}
	return res

}
