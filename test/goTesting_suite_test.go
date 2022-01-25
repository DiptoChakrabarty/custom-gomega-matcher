package test

import (
	"goTest/human"
	"log"
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
		AfterEach(func() {
			log.Println("This is child")
		})
		It("returns true", func() {
			person := human.Person{Age: 9, FirstName: "Pinku", LastName: "Das"}

			response := person.IsChild()

			Expect(response).To(BeTrue())
		})
	})

	Context("When Person is not child", func() {
		BeforeEach(func() {
			log.Println("not child instance")
		})
		It("returns false", func() {
			person := human.Person{Age: 35}

			response := person.IsChild()

			Expect(response).To(BeFalse())
		})
	})
})
