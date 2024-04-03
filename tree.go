package algorithm_utils

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NewTreeNode create a new tree node
func NewTreeNode(val int, left, right *TreeNode) *TreeNode {
	return &TreeNode{val, left, right}
}

// PreOrder pre-order traversal
func PreOrder(root *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}
	return append(append([]*TreeNode{root}, PreOrder(root.Left)...), PreOrder(root.Right)...)
}

// InOrder in-order traversal
func InOrder(root *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}
	return append(append(InOrder(root.Left), root), InOrder(root.Right)...)
}

// PostOrder post-order traversal
func PostOrder(root *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}
	return append(append(PostOrder(root.Left), PostOrder(root.Right)...), root)
}

// LevelOrder level-order traversal
func LevelOrder(root *TreeNode) (result []*TreeNode) {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return result
}
