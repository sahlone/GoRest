package functional_tests

import (
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/auth"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Authentication tests", func() {

	Describe("Pass::Authentication ", func() {

		Context("correct credentials", func() {
			It("should return an IDToken", func() {
				token, resp := GetIDToken("admin", "password")
				Expect(resp.StatusCode).To(Equal(200))
				Expect(token).To(Not(Equal("")))
				err := auth.ValidateToken(token)
				if err != nil {
					Fail("Error should be nil")
				}
			})

		})
	})

	Describe("Fail::Authentication ", func() {

		Context("invalid credentials", func() {
			It("should return 401 ", func() {
				_, resp := GetIDToken("false", "false")
				Expect(resp.StatusCode).To(Equal(401))
			})
		})
	})
})
