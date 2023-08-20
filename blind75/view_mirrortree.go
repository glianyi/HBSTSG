package blind

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }
//           2
//       1       3
//    2    5   6    7
//
//

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := mirrorTree(root.Right)
	right := mirrorTree(root.Left)

	root.Left = right
	root.Right = left
	return root
}

func mirrorTreeIter(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	list := make([]*TreeNode, 0)
	list = append(list, root)

	for len(list) > 0 {
		length := len(list)
		for i := 0; i < length; i++ {
			elem := list[i]
			if elem.Right != nil {
				list = append(list, elem.Right)
			}
			if elem.Left != nil {
				list = append(list, elem.Left)
			}
		}

		list = list[length:]
	}

	return root
}

// error 一个节点可能面临多次mirror 必须是回溯
func mirrorTreeL(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp
	mirrorTreeL(root.Left)
	mirrorTreeL(root.Right)

	return root
}

// error 一个节点可能面临多次mirror
func mirrorTreeE(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	root.Right = mirrorTreeL(root.Left)
	root.Left = (root.Right)

	return root
}
