package hash_table_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHashTable(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "HashTable Suite")
}
