package 热题100

/**
 * @Author: yirufeng
 * @Date: 2021/10/21 4:09 下午
 * @Desc:
 **/
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	//如果有一个为空，返回另外一个
	if root1 == nil {
		return root2
	}

	if root2 == nil {
		return root1
	}

	//此时说明两个都不为空
	root1.Left, root1.Right = mergeTrees(root1.Left, root2.Left), mergeTrees(root1.Right, root2.Right)
	root1.Val += root2.Val
	return root1
}
