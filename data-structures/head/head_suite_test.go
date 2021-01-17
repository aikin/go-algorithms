package head_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHead(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Head Suite")
}
