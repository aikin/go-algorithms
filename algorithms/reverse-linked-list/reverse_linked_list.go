package reverse_linked_list

import (
	. "github.com/aikin/go-algorithms/data-structures/linked-list"
	"fmt"
)

/*
Given 1->2->3->4->5->NULL

Because reversed linked list
return  5->4->3->2->1->NULL

复杂度分析：
 - 时间复杂度：O(n)
 - 空间复杂度：O(1)
*/
func ReverseLinkedList(head *Node) *Node {
	var prev *Node
	curr := head

	for curr != nil {
		temp := curr.Next

		curr.Next = prev
		prev = curr

		curr = temp
	}

	return prev
}


/*
Given 1->2->3->4->5->NULL

Because reversed linked list
return  5->4->3->2->1->NULL

复杂度分析：
 - 时间复杂度：O(n)
 - 空间复杂度：O(n) 由于使用递归，将会使用隐式栈空间。递归深度可能会达到 n 层。
*/
func ReverseLinkedListByRecursive(head *Node) *Node {
	if (head == nil || head.Next == nil) {
		fmt.Printf("%v------x>", head.Value)
		fmt.Println()
		return head
	}

	// fmt.Printf("%v->", head.Value)
	// fmt.Println()

	curr := ReverseLinkedListByRecursive(head.Next)

	fmt.Printf("%v---->", head.Value)
	fmt.Println()

	fmt.Printf("%v----->", head.Next.Next)
	fmt.Println()

	fmt.Printf("%v----->", head.Next.Value)
	fmt.Println()

	// 反转节点指向
	head.Next.Next = head
	head.Next = nil

	return curr
}