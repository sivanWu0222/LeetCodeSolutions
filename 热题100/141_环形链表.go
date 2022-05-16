package 热题100

/**
 * @Author: yirufeng
 * @Date: 2021/10/20 11:06 上午
 * @Desc:
 **/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		fast, slow = fast.Next.Next, slow.Next
		if fast == slow {
			return true
		}
	}

	return false
}
