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
		targetAge   int
		humanTarget Person
		want        bool
	}{
		{
			name:        "Compare Age Correct",
			targetAge:   10,
			humanTarget: Person{age: 10},
			want:        true,
		},
		{
			name:        "Compare Age Incorrect",
			targetAge:   11,
			humanTarget: Person{age: 12},
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
	age int
}

func (p *Person) Age() int {
	return p.age
}

func MatchAge(a int) types.GomegaMatcher {
	//fmt.Println(s)
	return &Person{age: a}
}

func (p *Person) Match(actual interface{}) (bool, error) {
	//fmt.Println("This is actuval", actual)
	//fmt.Println("This is person", p)
	pr := actual.(Person)
	if p.age != pr.age {
		return false, fmt.Errorf("Wrong Person")
	}
	return true, nil
}

func (p *Person) FailureMessage(actual interface{}) string {
	return fmt.Sprintf("Expected name to be %v but received %v", actual, p.age)
}

func (p *Person) NegatedFailureMessage(actual interface{}) string {
	return fmt.Sprintf("Expected name to be %v but received %v", actual, p.age)
}
