package main

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAdd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Add Function Suite")
}

var _ = Describe("add", func() {
	Context("with two integers", func() {
		It("should return their sum", func() {
			Expect(add(2, 3)).To(Equal(5))
			Expect(add(-1, 1)).To(Equal(0))
			Expect(add(0, 0)).To(Equal(0))
		})
	})
})
