package functional_tests

import (
	"fmt"
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/square/go-jose.v2/json"
)

var _ = Describe("List Recipe tests", func() {

	Describe("Pass::ListRecipe ", func() {

		for i := 0; i < 20; i++ {
			CreateRecipe(testRecipe, true)
		}
		Context("limit is defined", func() {
			It("should return 200 with correct no of record", func() {
				resp := ExecuteRequest("GET", "/recipes?limit=20", "", false)
				recipes := []model.Recipe{}
				json.NewDecoder(resp.Body).Decode(&recipes)
				Expect(resp.StatusCode).To(Equal(200))
				Expect(len(recipes)).To(Equal(20))
			})
		})
		Context("limit is not defined defined", func() {
			It("should return 200 with correct 10 no of record", func() {

				resp := ExecuteRequest("GET", "/recipes", "", false)
				recipes := []model.Recipe{}
				json.NewDecoder(resp.Body).Decode(&recipes)
				Expect(resp.StatusCode).To(Equal(200))
				Expect(len(recipes)).To(Equal(10))
			})
		})
		Context("limit & last record is defined ", func() {

			It("should return 200 with correct pagination", func() {

				for i := 0; i < 45; i++ {
					CreateRecipe(testRecipe, true)
				}
				lastRecords := []string{}
				lastRec := ""
				for i := 1; i <= 5; i++ {
					path := fmt.Sprintf("/recipes?limit=9&lastRecord=%v", lastRec)
					resp := ExecuteRequest("GET", path, "", false)
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
						lastRec = recipes[j].ID
					}
				}
			})
		})
	})

	Describe("Fail::GetRecipe", func() {

		Context("with incorrect limit", func() {
			It("should return 400 ", func() {
				resp := ExecuteRequest("GET", "/recipes?limit=101", "", false)
				Expect(resp.StatusCode).To(Equal(400))
			})
		})
	})
})

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}
