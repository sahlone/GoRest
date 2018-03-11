package functional_tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Delete Recipe", func() {

	Describe("Pass::DeleteRecipe ", func() {

		Context("correctly defined input", func() {

			It("should return 204 and archive flag set true", func() {
				result, resp := CreateRecipe(testRecipe, true)
				resp = DeleteRecipe(result.ID, true)
				Expect(resp.StatusCode).To(Equal(204))
				result, resp = GetRecipe(result.ID)
				Expect(resp.StatusCode).To(Equal(200))
				Expect(result.Archive).To(Equal(true))
			})
		})
	})

	Describe("Fail::DeleteRecipe ", func() {

		Context("correctly defined input with bad authentication", func() {

			It("should return 401", func() {
				result, _ := CreateRecipe(testRecipe, true)
				resp := DeleteRecipe(result.ID, false)
				Expect(resp.StatusCode).To(Equal(401))
			})
		})

		Context("incorrect recipe-id with correct authentication", func() {

			It("should return 404", func() {
				resp := DeleteRecipe("dummy-id", true)
				Expect(resp.StatusCode).To(Equal(404))
			})
		})

	})
})
