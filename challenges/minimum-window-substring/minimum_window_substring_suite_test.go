package minimum_window_substring_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMinimumWindowSubstring(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MinimumWindowSubstring Suite")
}
