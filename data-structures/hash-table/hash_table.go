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

func(ht *HashTable) HashCode(key string) int {
	hash := int32(0)
	for i := 0; i < len(key); i++ {
		hash = int32(hash<<5-hash) + int32(key[i])
		hash &= hash
	}
	return int(math.Abs(float64(hash)))
}

func(ht *HashTable) Position(key string) int {
	return ht.HashCode(key) % ht.Capacity
}

func(ht *HashTable) Put(key string, value string) {
	index := ht.Position(key)
	if ht.Table[index] == nil {
		ht.Table[index] = list.NewLinkedList()
	}

	item := &item{key: key, value: value}

	ht.Table[index].Append(item)

	ht.Size++ 
}


func(ht *HashTable) Get(key string) (interface{}) {
	index := ht.Position(key)

	l := ht.Table[index]

	var val *item
	for node := l.Head; node != nil; node = node.Next {
		if (node.Value.(*item).key == key) {
			val = node.Value.(*item)
		}
	}

	if val == nil {
		return nil
	}

	return val.value
}