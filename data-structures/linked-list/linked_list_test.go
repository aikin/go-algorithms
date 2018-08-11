package list

import (
	"testing"
	"fmt"
)

func TestLinkedList(t *testing.T) {
	l := NewList()
	fmt.Println(l)
	if l.Length != 0 {
		t.Fail()
	}
}
