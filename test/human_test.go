package test

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
)

func TestHumanStruct(t *testing.T) {
	tests := []struct {
		name      string
		humanTest Person
		want      bool
	}{
		{
			name:      "Pinku",
			humanTest: Person{name: "Pinku"},
			want:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewWithT(t)
			p1 := Person{name: "Pinku"}
			g.Expect(p1).To(ContainElement(MatchFirstName(tt.name)))
		})
	}
}

type Person struct {
	name string
}

func MatchFirstName(s string) types.GomegaMatcher {
	return &Person{s}
}

func (p *Person) Match(actual string) (bool, error) {
	if p.name != actual {
		return false, fmt.Errorf("Wrong Person")
	}
	return true, nil
}
