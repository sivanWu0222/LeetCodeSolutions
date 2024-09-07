package main

/**
 * @Author: yirufeng
 * @Date: 2021/10/10 10:09 上午
 * @Desc:
 **/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	//如果两个都为空
	if p == nil && q == nil {
		return true
	}

	//如果有一个为空
	if p == nil || q == nil {
		return false
	}

	//继续下一层的递归遍历
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
