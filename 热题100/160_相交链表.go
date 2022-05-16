package 热题100

/**
 * @Author: yirufeng
 * @Date: 2021/10/21 11:15 上午
 * @Desc: 得到两个链表各自的长度，先走一个长度差，然后再一起走
 **/

func getIntersectionNode(headA, headB *ListNode) *ListNode {

	cur1, cur2 := headA, headB
	length1, length2 := 0, 0

	//获取链表1的长度
	for cur1 != nil {
		cur1, length1 = cur1.Next, length1+1
	}

	//获取链表2的长度
	for cur2 != nil {
		cur2, length2 = cur2.Next, length2+1
	}

	cur1, cur2 = headA, headB

	//如果链表1长，先走length1-length2步
	if length1 > length2 {
		for i := 0; i < length1-length2; i++ {
			cur1 = cur1.Next
		}
	} else { //如果链表2长，先走length2-length1步
		for i := 0; i < length2-length1; i++ {
			cur2 = cur2.Next
		}
	}

	//此时就可以一起走
	for cur1 != cur2 {
		cur1, cur2 = cur1.Next, cur2.Next
		if cur1 == nil || cur2 == nil {
			return nil
		}
	}

	return cur1
}
