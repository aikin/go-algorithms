package merge_intervals_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMergeIntervals(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MergeIntervals Suite")
}
