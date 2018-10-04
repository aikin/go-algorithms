package doubly_linked_list


type DoublyLinkedList struct {
	Length int
	Head   *DoublyLinkedListNode
	Tail   *DoublyLinkedListNode
}

type DoublyLinkedListNode struct {
	Value interface {}
	Next *DoublyLinkedListNode
	Previous *DoublyLinkedListNode
}


func NewDoublyLinkedListNode(value interface{}) *DoublyLinkedListNode {
	return &DoublyLinkedListNode{Value: value}
}

func NewDoublyLinkedList() *DoublyLinkedList {
	l := new(DoublyLinkedList)
	l.Length = 0
	return l
}

func (l *DoublyLinkedList) Len() int {
	return l.Length
}

func (l *DoublyLinkedList) Append(value interface{}) {
	node := NewDoublyLinkedListNode(value)

	if l.Len() == 0 {
		l.Head = node
		l.Tail = l.Head
	} else {
		l.Tail.Next = node
		node.Previous = l.Tail
		l.Tail = node
	}

	l.Length++
}

func (l *DoublyLinkedList) Prepend(value interface{}) {
	node := NewDoublyLinkedListNode(value)

	if l.Len() == 0 {
		l.Head = node
		l.Tail = l.Head
	} else {
		l.Head.Previous = node
		node.Next = l.Head
		l.Head = node
	}

	l.Length++
}


func (l *DoublyLinkedList) Remove(value interface{}) bool {
	originalLength := l.Len()

	currentNode := l.Head
	
	for currentNode != nil {
		if currentNode.Value == value {
			removedNode := currentNode

			if removedNode == l.Head {
				l.Head = removedNode.Next
				if l.Head != nil {
					l.Head.Previous = nil
				}

				if removedNode == l.Tail {
					l.Tail = nil
				}
			} else if removedNode == l.Tail {
				l.Tail = removedNode.Previous
				l.Tail.Next = nil
			} else {
				currentNode.Next.Previous, currentNode.Previous.Next = currentNode.Previous, currentNode.Next				
			}

			l.Length--
		} 
		currentNode = currentNode.Next
	}

	return originalLength != l.Len()
}