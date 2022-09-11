/**
 @Author       cvenwu
 @Datetime     2022/9/10 10:12
 @Description
**/
package main

import "fmt"

// 树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//
// trimBST
// @Description: 看了宫水三叶的题解之后，按照递归的思路进行解决
// @param root: 根节点的值
// @param low: 最小边界
// @param high: 最大边界
// @return *TreeNode: 裁剪之后的树的根节点
//
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return root
	} else if root.Val < low { // 如果当前节点小于最小值，遍历到右子树
		return trimBST(root.Right, low, high)
	} else if root.Val > high { // 如果当前节点的值大于最大值，说明裁剪之后的树位于当前节点的左子树
		return trimBST(root.Left, low, high)
	}

	// 走到这里说明裁剪之后的树里面一定有当前节点，所以再裁剪左右子树，并返回
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}

// 使用迭代来
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	// 返回的结果的树的根节点
	retRoot := root
	for root != nil {
		// 如果当前节点的值大于最小值，左子树去掉

		// 如果当前节点的值小于
	}
}

func main() {
	fmt.Println("2022")
}
