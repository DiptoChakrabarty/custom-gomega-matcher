package test

import (
	"fmt"
	"testing"

	//	. "github.com/onsi/ginkgo/v2"

	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
)

/*
type HumanBeing interface {
	Name() string
}*/

func TestHumanStruct(t *testing.T) {
	tests := []struct {
		name       string
		targetName string
		humanTest  Person
		want       bool
	}{
		{
			name:       "Compare first names",
			targetName: "Pinku",
			humanTest:  Person{name: "Pinku"},
			want:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewWithT(t)
			p1 := Person{name: "Pinku"}
			g.Expect(p1).To(Equal(MatchFirstName(tt.targetName)))
		})
	}
}

type Person struct {
	name string
}

func (p *Person) Name() string {
	return p.name
}

func MatchFirstName(s string) types.GomegaMatcher {
	//fmt.Println(s)
	return &Person{name: s}
}

func (p *Person) Match(actual interface{}) (bool, error) {
	//fmt.Println("This is actuval", actual)
	//fmt.Println("This is person", p)
	pr := actual.(Person)
	if p.name != pr.name {
		return false, fmt.Errorf("Wrong Person")
	}
	return true, nil
}

func (p *Person) FailureMessage(actual interface{}) string {
	return fmt.Sprintf("Expected name to be %v but received %v", actual, p.name)
}

func (p *Person) NegatedFailureMessage(actual interface{}) string {
	return fmt.Sprintf("Expected name to be %v but received %v", actual, p.name)
}
