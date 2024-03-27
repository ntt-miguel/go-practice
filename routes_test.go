package routes_test

import (
	"bytes"
	"net/http"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Routes", func() {
	Describe("GET /fundings", func() {
		Context("when trying to get all fundings", func() {
			It("should return status 404", func() {
				requestBody := ""
				request, _ := http.NewRequest("GET", "http://localhost:8080/noroute", bytes.NewBuffer([]byte(requestBody)))
				request.Header.Set("Content-Type", "application/json")

				client := &http.Client{}
				response, err := client.Do(request)

				Expect(err).NotTo(HaveOccurred())
				Expect(response.StatusCode).To(Equal(http.StatusNotFound))
			})
			It("should return all the fundings", func() {
				requestBody := ""
				request, _ := http.NewRequest("GET", "http://localhost:8080/fundings", bytes.NewBuffer([]byte(requestBody)))
				request.Header.Set("Content-Type", "application/json")

				client := &http.Client{}
				response, err := client.Do(request)

				Expect(err).NotTo(HaveOccurred())
				Expect(response.StatusCode).To(Equal(http.StatusOK))
			})
		})
	})

})
