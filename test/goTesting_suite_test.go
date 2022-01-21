package test

import (
	"goTest/human"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGoTesting(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoTesting Suite")
}

var _ = Describe("Person is a Child", func() {
	Context("When Person is a child", func() {
		It("returns true", func() {
			person := human.Person{Age: 10}

			response := person.IsChild()

			Expect(response).To(BeFalse())
		})
	})
})
