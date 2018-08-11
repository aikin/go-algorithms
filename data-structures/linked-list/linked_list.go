package list

type List struct {
	Length int
	Head   *Node
	Tail   *Node
}

type Node struct {
	Value interface{}
	Prev  *Node
	Next  *Node
}

func NewNode(value interface{}) *Node {
	return &Node{Value: value}
}

func NewList() *List {
	l := new(List)
	l.Length = 0
	return l
}