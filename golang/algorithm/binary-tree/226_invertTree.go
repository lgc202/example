package binary_tree

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
