package 热题100

/**
 * @Author: yirufeng
 * @Date: 2021/10/24 11:50 上午
 * @Desc:
 **/

//递归
//func inorderTraversal(root *TreeNode) []int {
//	if root == nil {
//		return nil
//	}
//	//首先遍历左子树
//	ret := inorderTraversal(root.Left)
//	//遍历当前根节点
//	ret = append(ret, root.Val)
//	//遍历右子树
//	ret = append(ret, inorderTraversal(root.Right)...)
//	//返回结果
//	return ret
//}

//非递归
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	cur := root
	stack := []*TreeNode{root}
	var ret []int
	for len(stack) != 0 {
		for cur.Left != nil {
			cur = cur.Left
			stack = append(stack, cur)
		}
		//将当前节点的值加入到当前结果中
		node := stack[len(stack)-1]
		ret = append(ret, node.Val)
		stack = stack[:len(stack)-1]
		//如果当前节点的右子树不空，则移动到当前节点的右节点
		if node.Right != nil {
			cur = node.Right
			stack = append(stack, cur)
		}
	}

	return ret
}
