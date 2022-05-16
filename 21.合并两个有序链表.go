package main

/**
 * @Author: yirufeng
 * @Date: 2022/5/5 8:24 下午
 * @Email: yirufeng@foxmail.com
 * @GitHub: https://www.github.com/sivanWu0222
 * @Desc:
 **/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	dummyNode := &ListNode{}
	cur := dummyNode


	//如果两个链表都有值
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			//不能写成cur.Next, cur, list1 = list1, cur.Next, list1.Next，因为cur.Next还是指向nil呢，所以cur=cur.Next将会出错
			cur.Next, cur, list1 = list1, list1, list1.Next
		} else {
			cur.Next, cur, list2 = list2, list2, list2.Next
		}
	}

	//此时只剩下一个链表有值
	for list1 != nil {
		cur.Next, cur, list1 = list1, list1, list1.Next
	}

	for list2 != nil {
		cur.Next, cur, list2 = list2, list2, list2.Next
	}

	return dummyNode.Next
}