# goTesting
This is to Learn inDepth of Gomega and Ginkgo

Aim : Learn Custom Gomega Matcher

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

The human_test.go file contains a custom gomega matcher which can be used for testing
```sh
type Person struct {
	age []int
}
```

- This struct age attribute is tested
- The tests follow the format followed by kubernetes sub projects which I have observed
