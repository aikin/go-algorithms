package shell_sort_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestShellSort(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ShellSort Suite")
}
