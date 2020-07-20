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

		It("should generate proper index for specified keys", func() {
			hashTable := NewHashTable(32)

			Ω(hashTable.Position("a")).Should(Equal(1))
			Ω(hashTable.Position("akin")).Should(Equal(15))
			Ω(hashTable.Position("hash")).Should(Equal(14))
			Ω(hashTable.Position("index")).Should(Equal(18))
		})
	})

	Context("when get item to hashtable", func() {
		It("should can get item", func() {
			hashTable := NewHashTable(32)

			hashTable.Put("a", "shy")
			hashTable.Put("b", "boy")

		    value := hashTable.Get("a")

			Ω(value).Should(Equal("shy"))
		})

		It("should can get correct item when hash collisions", func() {
			hashTable := NewHashTable(32)

			hashTable.Put("a", "shy")
			hashTable.Put("b", "boy")
			hashTable.Put("a", "shy-new")

		    value := hashTable.Get("a")

			Ω(value).Should(Equal("shy-new"))
		})
	})

	Context("when put item to hashtabe", func() {
		It("should can put item", func() {
			hashTable := NewHashTable(32)

			hashTable.Put("kin", "shy")
			hashTable.Put("boy", "shy-new")

			hashTable.Get("a")

			Ω(hashTable.Size).Should(Equal(2))
			Ω(hashTable.Get("a")).Should(BeNil())
			Ω(hashTable.Get("kin")).Should(Equal("shy"))
		})
	})

})
