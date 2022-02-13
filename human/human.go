package human

import "fmt"

// Person Struct to Compare
type Person struct {
	Age []int
}

// Match expects the actual item which is compared to the
// target returned from the custom Gomega function
func (p *Person) Match(actual interface{}) (bool, error) {
	switch actual := actual.(type) {
	case Person:
		for i, j := range actual.Age {
			if j != p.Age[i] {
				return false, fmt.Errorf("Wrong Person")
			}
		}
	}
	return true, nil
}

// FailureMessage method for Person struct
func (p *Person) FailureMessage(actual interface{}) string {
	return fmt.Sprintf("Expected age to be %v but received %v", actual, p.Age)
}

// NegatedFailureMessage method for Person struct
func (p *Person) NegatedFailureMessage(actual interface{}) string {
	return fmt.Sprintf("Expected age to be %v but received %v", actual, p.Age)
}
