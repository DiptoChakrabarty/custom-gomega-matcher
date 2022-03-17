package test

import (
	"testing"

	"goTest/human"

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
			humanTarget: human.Person{Age: []int{30, 50}},
			want:        true,
		},
		{
			name:        "Compare Age Incorrect",
			targetAge:   []int{11, 25},
			humanTarget: human.Person{Age: []int{32, 60}},
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

// MatchAge is the Custom Gomega Testing Function
func MatchAge(a []int) types.GomegaMatcher {
	return &human.Person{Age: a}
}
