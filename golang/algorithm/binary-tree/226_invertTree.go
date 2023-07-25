package binary_tree

// invertTreeV1 利用递归函数的定义: 输入是树的根节点, 输出是一颗反转后的二叉树
// 所以分别把左节点和右节点作为输入得到两颗反转后的子树, 再把两颗子树交互位置即得
// 到反转后的二叉树
func invertTreeV1(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	left := invertTreeV1(root.Left)
	right := invertTreeV1(root.Right)

	root.Left = right
	root.Right = left

	return root
}

func invertTreeV2(root *TreeNode) *TreeNode {
	traverse(root)
	return root
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}

	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp

	traverse(root.Left)
	traverse(root.Right)
}
