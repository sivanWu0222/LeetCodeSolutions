package 热题100

/**
 * @Author: yirufeng
 * @Date: 2021/10/19 9:16 下午
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(preorder) != len(inorder) {
		return nil
	}

	//找到根节点
	root := &TreeNode{
		Val: preorder[0],
	}

	//从中序遍历序列中找到根节点的位置
	leftIndex := 0
	for leftIndex = 0; leftIndex < len(inorder); leftIndex++ {
		if inorder[leftIndex] == root.Val {
			break
		}
	}

	root.Left, root.Right = buildTree(preorder[1:leftIndex+1], inorder[:leftIndex]),
		buildTree(preorder[leftIndex+1:], inorder[leftIndex+1:])

	return root
}
