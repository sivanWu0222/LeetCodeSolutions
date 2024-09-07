package 热题100

/**
 * @Author: yirufeng
 * @Date: 2021/10/19 10:04 上午
 * @Desc:
 **/

type ListNode struct {
	Val  int
	Next *ListNode
}
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	dummyNode := &ListNode{}
	cur := dummyNode
	val, carry := 0, 0
	for l1 != nil || l2 != nil {

		val = 0

		if l1 != nil {
			val, l1 = val+l1.Val, l1.Next
		}

		if l2 != nil {
			val, l2 = val+l2.Val, l2.Next
		}

		cur.Next = &ListNode{
			Val: (val + carry) % 10,
		}

		cur, carry = cur.Next, (val + carry)/10
	}


	if carry != 0 {
		cur.Next = &ListNode{
			Val: carry,
		}
	}

	return dummyNode.Next
}
