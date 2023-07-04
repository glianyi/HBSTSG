package blind

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 	2
// 	/ \
// 1   4
//    / \
//   3   5

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var helper func(*TreeNode, int, int) bool

	helper = func(root *TreeNode, low, high int) bool {
		if root == nil {
			return true
		}

		if root.Val < low {
			return false
		}

		if root.Val > high {
			return false
		}
		return helper(root.Left, low, root.Val) && helper(root.Right, root.Val, high)
	}

	return helper(root, math.MinInt32, math.MaxInt32)
}
