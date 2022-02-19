# Custom Gomega Matcher

This Repository has code to implement custom gomega matcher

As per documentation as found [here](https://onsi.github.io/gomega/#adding-your-own-matchers) we can develop custom gomega matchers which can extend the power of our testing

```sh
type GomegaMatcher interface {
    Match(actual interface{}) (success bool, err error)
    FailureMessage(actual interface{}) (message string)
    NegatedFailureMessage(actual interface{}) (message string)
}

```

Custom Gomega Matchers must fall under this interface , hence there has to be
three methods defined for our objects.

## Files Present

* human/human.go
    - This contains the Person struct which is to be tested
    - This struct age attribute is tested
    - The file also contains relevant methods to implement the gomega matcher

```sh
type Person struct {
	age []int
}

```

* human_test.go
    - The human_test.go file contains the implementation of the Person struct
    - The tests follow the format followed by kubernetes sub projects which I have observed

The tests follow a format of
```sh
struct {
		name        string
		targetAge   []int
		humanTarget human.Person
		want        bool
	}
```
- Here name defines name of the test
- targetAge is the values we give
- humanTarget is the Person struct we provide
- want identifies if true or false
