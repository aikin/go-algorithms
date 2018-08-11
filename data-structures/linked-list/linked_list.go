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

func LinkedListNode(value interface{}) *Node {
	return &Node{Value: value}
}

func LinkedList() *List {
	l := new(List)
	l.Length = 0
	return l
}