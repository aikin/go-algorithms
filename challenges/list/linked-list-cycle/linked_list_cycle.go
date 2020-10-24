package linked_list_cycle

import (
	. "github.com/aikin/go-algorithms/data-structures/linked-list"
	"fmt"
)

/*
Given head = [3,2,0,-4], pos = 1

Because 链表中有一个环，其尾部连接到第二个节点。
return  true

复杂度分析：
 - 时间复杂度：O(n)
 - 空间复杂度：O(n)
*/
func HasCycle(head *Node) bool {

	seen := map[*Node]struct{}{}

	for head != nil {
		if _, ok := seen[head]; ok {
			return true
		}

		seen[head] = struct{}{}
		head = head.Next

	}
	fmt.Printf("%v---->", head.Value)
	fmt.Printf("%v---->", head.Next.Value)
	fmt.Printf("%v---->", head.Next.Next.Value)
	fmt.Printf("%v---->", head.Next.Next.Next.Value)
	fmt.Printf("%v---->", head.Next.Next.Next.Next.Value)
	fmt.Printf("%v---->", head.Next.Next.Next.Next.Next.Value)
	fmt.Printf("%v---->", head.Next.Next.Next.Next.Next.Next.Value)
	fmt.Printf("%v---->", head.Next.Next.Next.Next.Next.Next.Next.Value)

	return false

}

/*
Given head = [3,2,0,-4], pos = 1

Because 链表中有一个环，其尾部连接到第二个节点。
return  true

复杂度分析：
 - 时间复杂度：O(n)
 - 空间复杂度：O(1)
*/
func HasCycleWithFloyd(head *Node) bool {

	if (head == nil || head.Next == nil) {
		return false
	}

	slow, fast := head, head.Next

	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}

		slow = slow.Next
		fast = fast.Next.Next // slow.Next
	}

	return true
}