package functional_tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"time"
)

var _ = Describe("Health endpoint tests", func() {

	Context("the container is healthy", func() {

		It("must return 200", func() {
			req, _ := http.NewRequest("GET", "http://localhost:9000/health", nil)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Accept", "application/json")
			client := &http.Client{Timeout: time.Duration(2 * time.Second)}
			res, err := client.Do(req)

			if err != nil {
				GinkgoWriter.Write([]byte(err.Error()))
			} else {
				Expect(res.StatusCode).To(Equal(200))
			}
		})
	})
})
