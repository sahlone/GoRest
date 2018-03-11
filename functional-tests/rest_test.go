package functional_tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestRecipeService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Recipe Service Suite")
}
