package hash_table

import (
	"github.com/aikin/go-algorithms/data-structures/linked-list"
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