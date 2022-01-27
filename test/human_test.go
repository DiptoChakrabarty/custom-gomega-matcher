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
		name        string
		targetAge   []int
		humanTarget Person
		want        bool
	}{
		{
			name:        "Compare Age Correct",
			targetAge:   []int{10, 20},
			humanTarget: Person{age: []int{10, 20}},
			want:        true,
		},
		{
			name:        "Compare Age Incorrect",
			targetAge:   []int{11, 21},
			humanTarget: Person{age: []int{10, 20}},
			want:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewWithT(t)
			g.Expect(tt.humanTarget).Should(MatchAge(tt.targetAge))
		})
	}
}

type Person struct {
	age []int
}

func (p *Person) Age() []int {
	return p.age
}

func MatchAge(a []int) types.GomegaMatcher {
	//fmt.Println(s)
	return &Person{age: a}
}

func (p *Person) Match(actual interface{}) (bool, error) {
	//fmt.Println("This is actuval", actual)
	//fmt.Println("This is person", p)
	switch actual := actual.(type) {
	case Person:
		for i, j := range actual.age {
			if j != p.age[i] {
				return false, fmt.Errorf("Wrong Person")
			}
		}
	}
	return true, nil
}

func (p *Person) FailureMessage(actual interface{}) string {
	return fmt.Sprintf("Expected age to be %v but received %v", actual, p.age)
}

func (p *Person) NegatedFailureMessage(actual interface{}) string {
	return fmt.Sprintf("Expected age to be %v but received %v", actual, p.age)
}
