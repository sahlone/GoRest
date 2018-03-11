package functional_tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"fmt"
	"gopkg.in/square/go-jose.v2/json"
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/model"
)

var _ = Describe("Update Recipe tests", func() {

	Describe("Pass::UpdateRecipe ", func() {

		Context("correctly defined input", func() {
			It("should return 201", func() {
				oldRecipe, resp := CreateRecipe(testRecipe, true)
				Expect(resp.StatusCode).To(Equal(201))
				oldRecipe.Name="changedName"
				oldRecipe.Prep.Time=2
				oldRecipe.Prep.Total=3
				oldRecipe.Difficulty=3
				oldRecipe.IsVeg=true
				params, _ := json.Marshal(oldRecipe)
				paramstr := string(params)
				path:=fmt.Sprintf("recipes/%s", oldRecipe.ID)
				resp=ExecuteRequest("PUT",path,paramstr,true)
				Expect(resp.StatusCode).To(Equal(204))
				updatedRecipe:=model.Recipe{}
				updatedRecipe,_=GetRecipe(oldRecipe.ID)
				Expect(updatedRecipe.ID).To(Equal(oldRecipe.ID))
				Expect(updatedRecipe.Name).To(Equal("changedName"))
				Expect(int(updatedRecipe.Difficulty)).To(Equal(3))
				Expect(updatedRecipe.IsVeg).To(Equal(true))
				Expect(updatedRecipe.Prep.Time).To(Equal(2))
				Expect(updatedRecipe.Prep.Total).To(Equal(3))
			})

		})
	})

	Describe("Fail::CreateRecipe ", func() {

		Context("passed bad content", func() {
			It("should return 400 ", func() {
				oldRecipe, resp := CreateRecipe(testRecipe, true)
				path:=fmt.Sprintf("recipes/%s", oldRecipe.ID)
				resp=ExecuteRequest("PUT",path,"",true)
				Expect(resp.StatusCode).To(Equal(400))
			})
		})

		Context("passed incorrect recipeId", func() {
			It("should return 404 ", func() {
				oldRecipe, resp := CreateRecipe(testRecipe, true)
				params, _ := json.Marshal(oldRecipe)
				paramstr := string(params)
				path:=fmt.Sprintf("recipes/%s", "not-found-id")
				resp=ExecuteRequest("PUT",path,paramstr,true)
				Expect(resp.StatusCode).To(Equal(404))
			})
		})

		Context("passed correctly defined input with bad authentication", func() {
			It("should return 401 ", func() {
				oldRecipe, resp := CreateRecipe(testRecipe, true)
				params, _ := json.Marshal(oldRecipe)
				paramstr := string(params)
				path:=fmt.Sprintf("recipes/%s", oldRecipe.ID)
				resp=ExecuteRequest("PUT",path,paramstr,false)
				Expect(resp.StatusCode).To(Equal(401))
			})
		})
	})
})
