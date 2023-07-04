package blind

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}

	if subRoot == nil {
		return true
	}

	var help func(root *TreeNode) bool

	help = func(root *TreeNode) bool {
		if root == nil {
			return false
		}
		if isSameTree(root, subRoot) {
			return true
		}
		help(root.Left)
		help(root.Right)
		return false
	}

	return help(root)
}
