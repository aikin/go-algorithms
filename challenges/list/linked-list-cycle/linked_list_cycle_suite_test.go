package linked_list_cycle_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLinkedListCycle(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "LinkedListCycle Suite")
}
