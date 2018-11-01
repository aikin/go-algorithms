package hash_table_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/aikin/go-algorithms/data-structures/hash-table"
)

var _ = Describe("HashTable", func() {

	Context("When create empty hash table", func() {
		It("should create hash table of certain size", func() {
			 hashTable := NewHashTable(1000)
			 
			 Ω(hashTable.Capacity).Should(Equal(1000))
		})

		It("should generate proper hash for specified keys", func() {
			hashTable := NewHashTable(32)
			 
			Ω(hashTable.Capacity).Should(Equal(32))
			Ω(hashTable.HashCode("a")).Should(Equal(97))
			Ω(hashTable.HashCode("b")).Should(Equal(98))
			Ω(hashTable.HashCode("hash")).Should(Equal(3195150))
		})

		It("should generate proper hash for specified keys", func() {
			hashTable := NewHashTable(32)
			 
			Ω(hashTable.Capacity).Should(Equal(32))
			Ω(hashTable.HashCode("a")).Should(Equal(97))
			Ω(hashTable.HashCode("b")).Should(Equal(98))
			Ω(hashTable.HashCode("hash")).Should(Equal(3195150))
		})
	})

})
