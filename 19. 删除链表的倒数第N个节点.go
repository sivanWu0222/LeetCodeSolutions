package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyNode := &ListNode{}
	dummyNode.Next = head

	//删除倒数第n个节点我们需要保持n+1的距离
	fast, slow := dummyNode, dummyNode
	for i := 0; i < n+1; i++ {
		//如果n越界，直接返回空
		if fast == nil {
			return dummyNode.Next
		}
		fast = fast.Next
	}

	for fast != nil {
		fast, slow = fast.Next, slow.Next
	}

	//此时slow指向倒数第n个节点的前1个节点
	slow.Next = slow.Next.Next
	return dummyNode.Next
}