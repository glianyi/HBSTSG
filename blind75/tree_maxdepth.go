package blind

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + maxInt(maxDepth(root.Left), maxDepth(root.Right))
}

func maxDepthIter(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	maxDepth := 1

	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			e := queue[i]
			if e.Left != nil {
				queue = append(queue, e.Left)
			}
			if e.Right != nil {
				queue = append(queue, e.Left)
			}
		}
		queue = queue[length:]
		maxDepth += 1
	}

	return maxDepth
}
