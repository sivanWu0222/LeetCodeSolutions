package main

/**
 * @Author: yirufeng
 * @Date: 2021/10/10 10:31 上午
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
func isSymmetric(root *TreeNode) bool {
	return isSymmetricCore(root, root)
}

func isSymmetricCore(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 == nil || root2 == nil {
		return false
	}

	return root1.Val == root2.Val && isSymmetricCore(root1.Left, root2.Right) && isSymmetricCore(root1.Right, root2.Left)
}
