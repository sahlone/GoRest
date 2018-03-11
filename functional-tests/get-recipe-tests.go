package functional_tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Get Recipe tests", func() {

	Describe("Pass::GetRecipe ", func() {

		Context("passed correctly defined input", func() {
			It("should return 200", func() {
				newRecipe, resp := CreateRecipe(testRecipe, true)
				if resp.StatusCode != 201 {
					Fail("Should have created recipe")
				}
				result, resp := GetRecipe(newRecipe.ID)
				Expect(resp.StatusCode).To(Equal(200))
				Expect(result.ID).To(Equal(newRecipe.ID))
				Expect(result.Name).To(Equal(newRecipe.Name))
				Expect(result.Difficulty).To(Equal(newRecipe.Difficulty))
				Expect(result.Prep).To(Equal(newRecipe.Prep))

			})
		})
	})

	Describe("Fail::GetRecipe", func() {

		Context("with incorrect recipe-id", func() {
			It("should return 404 ", func() {
				_, resp := GetRecipe("dummy-id")
				Expect(resp.StatusCode).To(Equal(404))
			})
		})
	})
})
