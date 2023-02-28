package binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

func preOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		res = append(res, node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)
	return res
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Left)
		res = append(res, node.Val)
		dfs(node.Right)
	}

	dfs(root)
	return res
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Left)
		dfs(node.Right)
		res = append(res, node.Val)
	}

	dfs(root)
	return res
}

func levelTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var (
		res   []int
		queue []*TreeNode
	)

	queue = append(queue, root)
	for len(queue) > 0 {
		// 当前层有多少个元素
		length := len(queue)
		for length > 0 {
			length--
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}

			res = append(res, queue[0].Val)
			queue = queue[1:]
		}
	}

	return res
}
