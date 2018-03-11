package functional_tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"fmt"
	"gopkg.in/square/go-jose.v2/json"
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/model"
)

var _ = Describe("Patch Recipe tests", func() {

	Describe("Pass::PatchRecipe ", func() {

		Context("correctly defined input", func() {
			It("should return 201", func() {
				recipe, resp := CreateRecipe(testRecipe, true)
				Expect(resp.StatusCode).To(Equal(201))
				patch:=model.RecipePatch{}
				patch.Name="changedName"
				patch.Prep.Time=2
				patch.Prep.Total=3
				patch.Difficulty=3
				patch.IsVeg="true"
				params, _ := json.Marshal(patch)
				paramstr := string(params)
				path:=fmt.Sprintf("recipes/%s", recipe.ID)
				resp=ExecuteRequest("PATCH",path,paramstr,true)
				Expect(resp.StatusCode).To(Equal(204))
				updatedRecipe:=model.Recipe{}
				updatedRecipe,_=GetRecipe(recipe.ID)
				Expect(updatedRecipe.ID).To(Equal(recipe.ID))
				Expect(updatedRecipe.Name).To(Equal("changedName"))
				Expect(int(updatedRecipe.Difficulty)).To(Equal(3))
				Expect(updatedRecipe.IsVeg).To(Equal(true))
				Expect(updatedRecipe.Prep.Time).To(Equal(2))
				Expect(updatedRecipe.Prep.Total).To(Equal(3))
			})

		})
	})

	Describe("Fail::PatchRecipe ", func() {

		Context("passed bad content", func() {
			It("should return 400 ", func() {
				oldRecipe, resp := CreateRecipe(testRecipe, true)
				path:=fmt.Sprintf("recipes/%s", oldRecipe.ID)
				resp=ExecuteRequest("PATCH",path,"",true)
				Expect(resp.StatusCode).To(Equal(400))
			})
		})

		Context("passed incorrect recipeId", func() {
			It("should return 404 ", func() {
				_, resp := CreateRecipe(testRecipe, true)
				patch:=model.RecipePatch{}
				patch.Name="changedName"
				patch.Prep.Time=2
				patch.Prep.Total=3
				patch.Difficulty=3
				patch.IsVeg="true"
				params, _ := json.Marshal(patch)
				paramstr := string(params)
				path:=fmt.Sprintf("recipes/%s", "not-found-id")
				resp=ExecuteRequest("PATCH",path,paramstr,true)
				Expect(resp.StatusCode).To(Equal(404))
			})
		})

		Context("passed correctly defined input with bad authentication", func() {
			It("should return 401 ", func() {
				oldRecipe, resp := CreateRecipe(testRecipe, true)
				params, _ := json.Marshal(oldRecipe)
				paramstr := string(params)
				path:=fmt.Sprintf("recipes/%s", oldRecipe.ID)
				resp=ExecuteRequest("PATCH",path,paramstr,false)
				Expect(resp.StatusCode).To(Equal(401))
			})
		})
	})
})
