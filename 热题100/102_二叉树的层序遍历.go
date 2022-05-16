package 热题100

/**
 * @Author: yirufeng
 * @Date: 2021/10/19 4:32 下午
 * @Desc:
 **/

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var ret [][]int
	queue := []*TreeNode{root}

	for len(queue) != 0 {

		length := len(queue)
		curLevelRet := make([]int, length)

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
