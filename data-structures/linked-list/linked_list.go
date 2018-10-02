package list

import "fmt"

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

func (l *List) Append(value interface{}) {
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

func (l *List) Prepend(value interface{}) {
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

func (l *List) Find(value interface{}) *Node {
	if l.Len() == 0 {
		return (*Node)(nil)
	}

	currentNode := l.Head

	for currentNode != nil {
		if value != nil && currentNode.Value == value {
			return currentNode;
		}
		currentNode = currentNode.Next
	}
	return (*Node)(nil)
}


func (l *List) Remove(value interface{}) *Node {
	removedNode := (*Node)(nil)
	if l.Len() == 0 {
		return removedNode
	}

	for l.Head != nil && l.Head.Value == value {
		removedNode = l.Head
		l.Head = l.Head.Next
		l.Length--
	}

	currentNode := l.Head
	if currentNode == nil {
		return removedNode
	}

	for currentNode.Next != nil {
		if currentNode.Next.Value == value {
			removedNode = currentNode.Next
			currentNode.Next = currentNode.Next.Next
			l.Length--
		} else {
			currentNode = currentNode.Next
		}
	}

	if l.Tail.Value == value {
		l.Tail = currentNode
	}

	return removedNode
}

func (l *List) Print() {
	node := l.Head
	for node != nil {
		fmt.Printf("%v->", node.Value)
		node = node.Next
	}
	fmt.Println()
}
