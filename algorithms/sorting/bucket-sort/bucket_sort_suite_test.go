package bucket_sort_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBucketSort(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BucketSort Suite")
}
