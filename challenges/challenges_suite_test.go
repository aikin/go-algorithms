package challenges_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestChallenges(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Challenges Suite")
}
