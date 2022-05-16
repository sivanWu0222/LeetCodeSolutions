package main

/**
 * @Author: yirufeng
 * @Date: 2021/10/10 10:11 上午
 * @Desc:

思路：遍历到当前节点的时候，
通过后序遍历获得左右子树的最小值和最大值，
如果当前节点比左子树最大大且比右子树最小小，那么说明是对的

 **/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	return isValidBSTCore(root).isValidBST
}

type Ret struct {
	min        int
	max        int
	isValidBST bool
}

func isValidBSTCore(root *TreeNode) *Ret {

	if root == nil {
		return &Ret{
			isValidBST: true,
		}
	}

	lRet, rRet := isValidBSTCore(root.Left), isValidBSTCore(root.Right)

	ret := &Ret{
		min: lRet.min,
		max: rRet.max,
	}

	//如果左右两个都有效，并且root的值位于(lRet.max, rRet.min)之间
	if lRet.isValidBST && rRet.isValidBST && root.Val > lRet.max && root.Val < rRet.min {
		ret.isValidBST = true
	}

	return ret
}
