package hash_table

import (
	"math"
	"errors"
	"github.com/aikin/go-algorithms/data-structures/linked-list"
	// "fmt"
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

	foundItem, err := ht.find(index, key)

	if err != nil {
		ht.Table[index].Append(item)
		ht.Size++ 		
	} else {
		foundItem.value = value
	}

}


func(ht *HashTable) Get(key string) (interface{}) {
	index := ht.Position(key)
	l := ht.Table[index]

	if l == nil {
		return nil
	}

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

func (ht *HashTable) find(i int, key string) (*item, error) {
	l := ht.Table[i]
	var found *item	

	for node := l.Head; node != nil; node = node.Next {
		 item := node.Value.(*item)
		if item.key == key {
			found = item
		}
	}
	if found == nil {
		return nil, errors.New("Not Found")
	}

	return found, nil
}