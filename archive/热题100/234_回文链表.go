package 热题100

/**
 * @Author: yirufeng
 * @Date: 2021/10/24 4:36 下午
 * @Desc:
 **/

func isPalindrome(head *ListNode) bool {
	fast, slow := head, head

	var prev *ListNode

	//快慢边走，边反转前面一部分
	for fast != nil && fast.Next != nil {
		fast, prev, slow.Next, slow = fast.Next, slow, prev, slow.Next
	}

	//记录没有反转部分的第一个节点
	temp := slow
	//此时prev指向slow的前一个节点,也就是我们反转之后的链表的头部
	//说明链表有奇数个节点
	if fast == nil {
		//此时慢指向中间那一个，因此让慢移动一步
		slow = slow.Next
	}

	//说明链表有偶数个节点
	//此时慢指向下一半链表的第一个
	//快指针恢复到链表开头
	fast = head

	//不断比较
	for slow != nil && fast != nil {
		if fast.Val != head.Val {
			return false
		}
		slow, fast = slow.Next, fast.Next
	}

	//将前部分反转的链表进行还原
	prev.Next=temp
	slow = prev
	for prev != nil {
		slow, prev = prev, prev.Next
	}



	return false
}
