package 热题100

/**
 * @Author: yirufeng
 * @Date: 2021/10/19 4:22 下午
 * @Desc:
 **/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	dummyNode := &ListNode{}
	cur := dummyNode
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next, cur, l1 = l1, l1, l1.Next
		} else {
			cur.Next, cur, l2 = l2, l2, l2.Next
		}
	}

	if l1 != nil {
		cur.Next = l1
	}

	if l2 != nil {
		cur.Next = l2
	}

	return dummyNode.Next
}
