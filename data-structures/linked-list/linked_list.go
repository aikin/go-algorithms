package list

type List struct {
	Length int
	Head   *Node
	Tail   *Node
}

type Node struct {
	Value interface{}
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

func (l *List) Len() int {
	 return l.Length
}

func (l *List) Append(value interface{})  {
	node := NewNode(value)

	if l.Len() == 0 {
		l.Head = node
		l.Tail = l.Head
	} else {
		l.Tail.Next = node
		l.Tail = node
	}
	l.Length++
}

func (l *List) Prepend(value interface{})  {
	node := NewNode(value)

	if l.Len() == 0 {
		l.Head = node
		l.Tail = l.Head
	} else {
		node.Next = l.Head
		l.Head = node
	}
	l.Length++
}
func (l *List) Remove(value interface{}) interface{} {
	return nil
}