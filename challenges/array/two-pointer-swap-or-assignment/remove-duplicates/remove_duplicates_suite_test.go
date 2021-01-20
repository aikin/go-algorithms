package remove_duplicates_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRemoveDuplicates(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RemoveDuplicates Suite")
}
