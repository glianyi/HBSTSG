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
func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32

	var maxPathFunc func(root *TreeNode) int

	maxPathFunc = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		maxLeftSum := maxInt(maxPathFunc(root.Left), 0)
		maxRightSum := maxInt(maxPathFunc(root.Right), 0)

		maxSum = maxInt(maxSum, root.Val+maxLeftSum+maxRightSum)

		return root.Val + maxInt(maxLeftSum, maxRightSum)

	}

	maxPathFunc(root)
	return maxSum
}

// 当前节点的最大路径： max(自己，自己+左边，自己+右边，自己 + 左边 + 右边）
// 	当前节点作为子节点时的贡献：max(自己，自己+左边，自己+右边）
// 	后者相对前者，少了左右都存在的情况。因为作为子节点时，一条路如果同时包含左右，根就被包含了2次，不符合题目只出现一次的限制了
