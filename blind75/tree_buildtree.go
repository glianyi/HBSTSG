package blind

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func indexOfElem(elems []int, v int) int {
	for i, e := range elems {
		if e == v {
			return i
		}
	}
	return -1
}

// preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{}
	root.Val = preorder[0]
	index := indexOfElem(inorder, root.Val)
	root.Left = buildTree(preorder[1:1+len(inorder[:index])], inorder[:index])
	root.Right = buildTree(preorder[1+len(inorder[:index]):], inorder[index+1:])

	return root
}
