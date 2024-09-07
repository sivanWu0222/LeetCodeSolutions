package main

/**
 * @Author: yirufeng
 * @Date: 2021/10/10 1:55 下午
 * @Desc:
 **/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	ret := 0

	for i := 0; i < len(root.Children); i++ {
		//TODO: 这里进行了优化，避免计算两次 maxDepth(root.Children[i]) 增大时间复杂度
		temp := maxDepth(root.Children[i])
		if ret < temp {
			ret = temp
		}
	}

	return ret + 1
}
