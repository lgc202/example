package binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

func preOrderTraverse(root *TreeNode) []int {
	return nil
}

func inorderTraversal(root *TreeNode) []int {
	return nil
}

func postorderTraversal(root *TreeNode) []int {
	return nil
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
