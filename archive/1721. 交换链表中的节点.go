/**
 @Author       cvenwu
 @Datetime     2022/9/10 22:49
 @Description
**/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodes(head *ListNode, k int) *ListNode {

	n := 0
	kthNode := head
	// 先走k-1步，记录第k个节点的值
	for i := 0; i < k-1; i++ {
		kthNode = kthNode.Next
		n++
	}
	cur := kthNode
	// 然后走到最后，同时记录链表长度n
	for cur != nil {
		cur = cur.Next
		n++
	}
	// 找到倒数第k个节点，也就是第n-k+1个节点，从第1个节点走n-k步到达该节点
	cur = head
	for i := 0; i < n-k; i++ {
		cur = cur.Next
	}

	cur.Val, kthNode.Val = kthNode.Val, cur.Val
	return head
}
