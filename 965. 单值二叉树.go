/**
 @Author       cvenwu
 @Datetime     2022/9/15 21:16
 @Description
**/
package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//
// isUnivalTree
// @Description: 比较当前根节点与左右子树的大小，然后递归到深层次遍历当前节点的左右子树
// @param root: 根节点
// @return bool: 是否为单值二叉树
//
func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	// 左子树有节点其不与根节点相同
	if root.Left != nil && root.Val != root.Left.Val {
		return false
	}

	// 右子树有节点且不与根节点相同
	if root.Right != nil && root.Val != root.Right.Val {
		return false
	}

	// 递归遍历左右子树
	return isUnivalTree(root.Left) && isUnivalTree(root.Right)
}
