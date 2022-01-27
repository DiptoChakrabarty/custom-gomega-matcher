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
		name      string
		targetAge int
		humanTest Person
		want      bool
	}{
		{
			name:      "Compare first names",
			targetAge: 10,
			humanTest: Person{age: 11},
			want:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewWithT(t)
			p1 := Person{age: 10}
			g.Expect(p1.Age()).To(Equal(MatchAge(tt.targetAge)))
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

func (p Person) Match(actual interface{}) (bool, error) {
	//fmt.Println("This is actuval", actual)
	//fmt.Println("This is person", p)
	pr := actual.(int)
	if p.age != pr {
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
