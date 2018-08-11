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

func NewLinkedList() *List {
	l := new(List)
	l.Length = 0
	return l
}

func (l *List) Append(value interface{})  {
	node := NewNode(value)

	if l.Length == 0 {
		l.Head = node
		l.Tail = l.Head
	}
	l.Length++
}