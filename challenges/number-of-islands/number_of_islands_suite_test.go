package number_of_islands_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestNumberOfIslands(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "NumberOfIslands Suite")
}
