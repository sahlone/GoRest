package functional_tests

import (
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/square/go-jose.v2/json"
	"github.com/google/uuid"
)

var _ = Describe("SearchRecipe tests", func() {

	Describe("Pass::SearchRecipe ", func() {

		Context("recipe with proper fields", func() {
			It("should return 200 when searching by name", func() {
				name:=uuid.New().String()
				recipe := model.Recipe{
					Name:       name,
					Prep:       model.Preparation{1, 1},
					Difficulty: 1,
					IsVeg:      true,
				}
				CreateRecipe(recipe, true)
				search := model.SearchCriteria{SearchByName: name}
				payload, _ := json.Marshal(search)
				resp := ExecuteRequest("POST", "recipes/search", string(payload), false)
				recipes := []model.Recipe{}
				json.NewDecoder(resp.Body).Decode(&recipes)
				Expect(resp.StatusCode).To(Equal(200))
				Expect(len(recipes)).To(Equal(1))
				Expect(recipes[0].Name).To(Equal(recipe.Name))
			})
		})
		Context("limit is not defined defined", func() {
			It("should return 200 with correct 10 no of record", func() {

				recipe := model.Recipe{
					Name:       "searchRecipe",
					Prep:       model.Preparation{1, 1},
					Difficulty: 1,
					IsVeg:      true,
				}
				for i:=0;i<12;i++ {
					CreateRecipe(recipe, true)
				}
				search := model.SearchCriteria{SearchByName: "searchRecipe"}
				payload, _ := json.Marshal(search)
				resp := ExecuteRequest("POST", "recipes/search", string(payload), false)
				recipes := []model.Recipe{}
				json.NewDecoder(resp.Body).Decode(&recipes)
				Expect(resp.StatusCode).To(Equal(200))
				Expect(len(recipes)).To(Equal(10))
				Expect(recipes[0].Name).To(Equal(recipe.Name))
			})
		})
		Context("limit & start record is defined ", func() {

			It("should return 200 with correct pagination", func() {

				recipe := model.Recipe{
					Name:       "t1",
					Prep:       model.Preparation{1, 1},
					Difficulty: 1,
					IsVeg:      true,
				}
				for i := 0; i < 45; i++ {
					CreateRecipe(recipe, true)
				}
				lastRecords := []string{}
				for i := 1; i <= 5; i++ {
					search := model.SearchCriteria{SearchByName: "t1",Start:(i-1)*9,Limit:9}
					payload, _ := json.Marshal(search)
					resp := ExecuteRequest("POST", "recipes/search", string(payload), false)
					recipes := []model.Recipe{}
					json.NewDecoder(resp.Body).Decode(&recipes)
					Expect(resp.StatusCode).To(Equal(200))
					Expect(len(recipes)).To(Equal(9))
					for j := 0; j < 9; j++ {
						is := contains(lastRecords, recipes[j].ID)
						Expect(is).To(Equal(false))
					}
					lastRecords = make([]string, 9)
					for j := 0; j < 9; j++ {
						lastRecords = append(lastRecords, recipes[j].ID)
					}
				}
			})
		})
	})

	Describe("Fail::SearchRecipe", func() {

		Context("with incorrect limit", func() {
			It("should return 400 ", func() {
				search := model.SearchCriteria{SearchByName: "t1",Start:0,Limit:102}
				payload, _ := json.Marshal(search)
				resp := ExecuteRequest("POST", "recipes/search", string(payload), false)
				Expect(resp.StatusCode).To(Equal(400))
			})
		})
	})
})

