package tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestGraphigo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Graphigo Test Suite")
}
