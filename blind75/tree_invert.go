package blind

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Right = left
	root.Left = right

	return root
}

func invertTreeIter(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	var parent *TreeNode

	queue := []*TreeNode{}
	queue = append(queue, root)

	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			parent = queue[i]
			parent.Left, parent.Right = parent.Right, parent.Left
			if parent.Left != nil {
				queue = append(queue, parent.Left)
			}
			if parent.Right != nil {
				queue = append(queue, parent.Right)
			}
		}

		queue = queue[length:]
	}

	return root
}
