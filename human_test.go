package test

import (
	"fmt"
	"testing"

	"github.com/DiptoChakrabarty/human"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
)

func TestHumanStruct(t *testing.T) {
	tests := []struct {
		name        string
		targetAge   []int
		humanTarget human.Person
		want        bool
	}{
		{
			name:        "Compare Age Correct",
			targetAge:   []int{10, 20},
			humanTarget: human.Person{age: []int{10, 20}},
			want:        true,
		},
		{
			name:        "Compare Age Incorrect",
			targetAge:   []int{11, 21},
			humanTarget: human.Person{age: []int{10, 20}},
			want:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewWithT(t)
			if tt.want {
				g.Expect(tt.humanTarget).Should(MatchAge(tt.targetAge))
			} else {
				g.Expect(tt.humanTarget).ShouldNot(MatchAge(tt.targetAge))
			}

		})
	}
}

// MatchAge is the Custom Gomega Testing Function
func MatchAge(a []int) types.GomegaMatcher {
	return &human.Person{age: a}
}

// Match expects the actual item which is compared to the
// target returned from the custom Gomega function
func (p *human.Person) Match(actual interface{}) (bool, error) {
	switch actual := actual.(type) {
	case human.Person:
		for i, j := range actual.age {
			if j != p.age[i] {
				return false, fmt.Errorf("Wrong Person")
			}
		}
	}
	return true, nil
}

func (p *human.Person) FailureMessage(actual interface{}) string {
	return fmt.Sprintf("Expected age to be %v but received %v", actual, p.age)
}

func (p *human.Person) NegatedFailureMessage(actual interface{}) string {
	return fmt.Sprintf("Expected age to be %v but received %v", actual, p.age)
}
