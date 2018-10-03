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

	if (l.Len() == 0) {
		l.Head = node
		l.Tail = l.Head
	}

	l.Length++
}