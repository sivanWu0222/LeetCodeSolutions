/*
 * @Author: cvenwu cvenwu@tencent.com
 * @Date: 2022-10-09 23:20:38
 * @LastEditors: cvenwu cvenwu@tencent.com
 * @LastEditTime: 2022-10-09 23:21:53
 * @FilePath: /LeetCodeSolutions/206. 反转链表.go
 * @Description:
 *
 * Copyright (c) 2022 by cvenwu cvenwu@tencent.com, All Rights Reserved.
 */
package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		prev, cur, cur.Next = cur, cur.Next, prev
	}
	return prev
}

func main() {

}
