package functional_tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Create Recipe tests", func() {

	Describe("Pass::CreateRecipe ", func() {

		Context("correctly defined input", func() {
			It("should return 201", func() {
				result, resp := CreateRecipe(testRecipe, true)
				Expect(resp.StatusCode).To(Equal(201))
				Expect(result.ID).To(Not(Equal("")))
				Expect(result.Name).To(Equal("Test"))
				Expect(int(result.Difficulty)).To(Equal(1))
				Expect(result.IsVeg).To(Equal(true))
			})

		})
	})

	Describe("Fail::CreateRecipe ", func() {

		Context("passed bad content", func() {
			It("should return 400 ", func() {
				res := ExecuteRequest("POST", "recipes", "", true)
				Expect(res.StatusCode).To(Equal(400))
			})
		})
		Context("passed correctly defined input with bad authentication", func() {
			It("should return 401 ", func() {
				_, resp := CreateRecipe(testRecipe, false)
				Expect(resp.StatusCode).To(Equal(401))
			})
		})

	})
})
