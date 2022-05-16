package main

/**
 * @Author: yirufeng
 * @Date: 2021/10/10 1:52 下午
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

//方法一：递归
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	}

	//根节点的值过大
	if root.Val > val {
		return searchBST(root.Left, val)
	}

	return searchBST(root.Right, val)
}


//方法二：循环
func searchBST(root *TreeNode, val int) *TreeNode {
	//如果不空且节点的值不等于val，就一直循环
	for root != nil && root.Val != val {
		if root.Val > val {
			root = root.Left
		} else {
			root = root.Right
		}
	}

	return root
}
