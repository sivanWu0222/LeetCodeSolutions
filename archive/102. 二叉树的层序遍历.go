package main

/**
 * @Author: yirufeng
 * @Date: 2021/10/10 10:34 上午
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
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	var ret [][]int
	for len(queue) != 0 {
		length := len(queue)
		curLevelRet := make([]int, len(queue))
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			curLevelRet[i] = node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ret = append(ret, curLevelRet)
	}

	return ret
}
