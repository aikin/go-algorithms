package priority_queue_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPriorityQueue(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PriorityQueue Suite")
}
