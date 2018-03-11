package functional_tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"fmt"
)

var _ = Describe("Rate Recipe tests", func() {

	Describe("Pass::RateRecipe ", func() {

		Context("correctly defined input", func() {
			It("should return 204", func() {
				result, resp := CreateRecipe(testRecipe, true)
				Expect(resp.StatusCode).To(Equal(201))
				path:=fmt.Sprintf("recipes/%s/rating/%d",result.ID,2)
				resp=ExecuteRequest("PUT",path,"",true)
				Expect(resp.StatusCode).To(Equal(204))
			})

		})
	})

	Describe("Fail::RateRecipe ", func() {

		Context("passed bad content", func() {
			It("should return 400 ", func() {
				result, resp := CreateRecipe(testRecipe, true)
				Expect(resp.StatusCode).To(Equal(201))
				path:=fmt.Sprintf("recipes/%s/rating/%d",result.ID,10)
				println(path)
				resp=ExecuteRequest("PUT",path,"",true)
				Expect(resp.StatusCode).To(Equal(400))
			})
		})
		Context("passed correctly defined input with bad authentication", func() {
			It("should return 401 ", func() {
				result, resp := CreateRecipe(testRecipe, true)
				path:=fmt.Sprintf("recipes/%s/rating/%d",result.ID,10)
				resp=ExecuteRequest("PUT",path,"",false)
				Expect(resp.StatusCode).To(Equal(401))
			})
		})
		Context("passed incorrect recipe id", func() {
			It("should return 400 ", func() {
				path:=fmt.Sprintf("recipes/%s/rating/%d","dummy-id",2)
				resp:=ExecuteRequest("PUT",path,"",true)
				Expect(resp.StatusCode).To(Equal(404))
			})
		})

	})
})
