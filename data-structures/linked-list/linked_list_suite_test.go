package list_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLinkedList(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "LinkedList Suite")
}
