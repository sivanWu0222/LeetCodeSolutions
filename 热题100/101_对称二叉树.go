package 热题100

/**
 * @Author: yirufeng
 * @Date: 2021/10/19 4:36 下午
 * @Desc:
 **/

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSymmetricCore(root.Left, root.Right)

}

func isSymmetricCore(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 == nil || root2 == nil {
		return false
	}

	return root1.Val == root2.Val && isSymmetricCore(root1.Left, root2.Right) && isSymmetricCore(root1.Right, root2.Left)
}
