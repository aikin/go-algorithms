package move_zeroes_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMoveZeroes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MoveZeroes Suite")
}
