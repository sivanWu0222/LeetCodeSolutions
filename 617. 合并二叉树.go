package main

/**
 * @Author: yirufeng
 * @Date: 2022/4/25 3:50 下午
 * @Email: yirufeng@foxmail.com
 * @GitHub: https://www.github.com/sivanWu0222
 * @Desc:
 **/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}



	if root2 == nil {
		return root1
	}

	//如果两个都不空
	root1.Val += root2.Val
	//如果左子树没有直接进行返回
	root1.Left, root1.Right = mergeTrees(root1.Left, root2.Left), mergeTrees(root1.Right, root2.Right)

	return root1
}