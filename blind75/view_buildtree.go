package blind

// 先root 中root 后root
// Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
// Output: [3,9,20,null,null,15,7]

func buildTreeView(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{}
	root.Val = preorder[0]
	ri := 0
	for i, v := range inorder {
		if root.Val == v {
			ri = i
		}
	}
	root.Left = buildTree(preorder[1:ri+1], inorder[:ri])
	root.Right = buildTree(preorder[ri+1:], inorder[ri+1:])

	return root
}
