package main

/**
 * @Author: yirufeng
 * @Date: 2021/10/15 8:16 上午
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
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}

	//层次遍历的时候每次遍历当前层的时候使用sum统计
	queue := []*TreeNode{root}

	//每一层的大小
	var length int
	//每一层的和
	var sum float64
	//结果数组
	ret := []float64{}

	for len(queue) != 0 {
		length = len(queue)
		sum = 0
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			//加入左节点
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			//加入右节点
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			sum += float64(node.Val)
		}
		ret = append(ret, sum/float64(length))
	}
	return ret
}
