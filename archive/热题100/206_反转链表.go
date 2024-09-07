package 热题100

/**
 * @Author: yirufeng
 * @Date: 2021/10/20 11:04 上午
 * @Desc:
 **/

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		prev, cur.Next, cur = cur, prev, cur.Next
	}
	return prev
}
