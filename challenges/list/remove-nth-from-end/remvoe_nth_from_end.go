package remove_nth_from_end

import (
	. "github.com/aikin/go-algorithms/data-structures/linked-list"
)

func getLength(head *Node) int {
	length := 0
	for ; head != nil; head = head.Next {
		length++
	}
	return length
}

/*
Given head = [1,2,3,4,5], n = 2

Because 删除倒数第二个元素
return  [1,2,3,5]

复杂度分析：
 - 时间复杂度：O(L)
 - 空间复杂度：O(1)
*/
func RemoveNthFromEnd(head *Node, n int) *Node {
	len := getLength(head)

	dummy := &Node{0, head}

	cur := dummy

	for i := 0; i < len - n; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummy.Next
}