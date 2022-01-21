package goTesting_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGoTesting(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoTesting Suite")
}
