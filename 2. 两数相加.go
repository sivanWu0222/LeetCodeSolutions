package main

/**
 * @Author: yirufeng
 * @Date: 2021/10/15 10:24 上午
 * @Desc:
 **/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//建立一个dummy node
	dummyNode := &ListNode{}
	cur := dummyNode

	val := 0
	carry := 0
	for l1 != nil || l2 != nil {
		val = 0
		if l1 != nil {
			l1, val = l1.Next, val+l1.Val
		}
		if l2 != nil {
			l2, val = l2.Next, val+l2.Val
		}
		cur.Next = &ListNode{
			Val: (val + carry) % 10,
		}
		carry, cur = (val+carry)/10, cur.Next
	}

	//到最后可能进位还有值
	if carry != 0 {
		cur.Next=&ListNode{
			Val: carry,
		}
	}

	return dummyNode.Next
}
