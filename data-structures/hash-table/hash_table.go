package hash_table

import (
	"github.com/aikin/go-algorithms/data-structures/linked-list"
	"math"
)

type HashTable struct {
	Table    map[int]*list.List
	Size     int
	Capacity int
}

type item struct {
	key   string
	value interface{}
}

func NewHashTable(cap int) *HashTable {
	table := make(map[int]*list.List, cap)
	
	return &HashTable{Table: table, Size: 0, Capacity: cap}
}

func(ht *HashTable) Hash(key string) int {
	hash := int32(0)
	for i := 0; i < len(key); i++ {
		hash = int32(hash<<5-hash) + int32(key[i])
		hash &= hash
	}
	return int(math.Abs(float64(hash)))
}