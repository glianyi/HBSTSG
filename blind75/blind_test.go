package blind

import (
	"fmt"
	"testing"
)

func makeTree() *TreeNode {
	// 	3
	// 	/ \
	// 9  20
	// 	 /  \
	// 	15   7
	l2l := TreeNode{
		Left:  nil,
		Right: nil,
		Val:   15,
	}

	l2r := TreeNode{
		Left:  nil,
		Right: nil,
		Val:   7,
	}

	l1l := TreeNode{
		Left:  nil,
		Right: nil,
		Val:   9,
	}

	l1r := TreeNode{
		Left:  &l2l,
		Right: &l2r,
		Val:   20,
	}

	r := TreeNode{
		Left:  &l1l,
		Right: &l1r,
		Val:   -10,
	}

	return &r
}

func makeAnotherTree() *TreeNode {
	// 	3
	// 	/ \
	// 9  20
	// 	 /  \
	// 	15   7
	l2l := TreeNode{
		Left:  nil,
		Right: nil,
		Val:   15,
	}

	l2r := TreeNode{
		Left:  nil,
		Right: nil,
		Val:   7,
	}

	l1l := TreeNode{
		Left:  nil,
		Right: nil,
		Val:   9,
	}

	l1r := TreeNode{
		Left:  &l2l,
		Right: &l2r,
		Val:   20,
	}

	r := TreeNode{
		Left:  &l1l,
		Right: &l1r,
		Val:   3,
	}

	return &r
}

func makeBST() *TreeNode {
	// 	2
	// 	/ \
	// 1   4
	//    / \
	//   3   5
	l2l := TreeNode{
		Left:  nil,
		Right: nil,
		Val:   0,
	}

	l2r := TreeNode{
		Left:  nil,
		Right: nil,
		Val:   5,
	}

	l1l := TreeNode{
		Left:  nil,
		Right: nil,
		Val:   1,
	}

	l1r := TreeNode{
		Left:  &l2l,
		Right: &l2r,
		Val:   4,
	}

	r := TreeNode{
		Left:  &l1l,
		Right: &l1r,
		Val:   2,
	}

	return &r
}

func printTree(r *TreeNode) {
	if r == nil {
		return
	}
	print(r.Val, ",")
	printTree(r.Left)
	printTree(r.Right)
}

func TestMaxDepth(t *testing.T) {
	root := makeTree()

	println(maxDepth(root))
	println(maxDepthIter(root))
}

func TestIsSame(t *testing.T) {
	root := makeTree()
	o := makeAnotherTree()
	println(isSameTree(root, root))
	println(isSameTree(root, nil))
	println(isSameTree(nil, root))
	println(isSameTree(root, o))
}

func TestInvert(t *testing.T) {
	// 	3,9,20,15,7,
	// 3,20,7,15,9,
	r := makeTree()
	printTree(r)
	println()
	// i := invertTree(r)
	i := invertTreeIter(r)
	printTree(i)
	println()
}

func TestPathSum(t *testing.T) {
	r := makeTree()
	printTree(r)
	println()
	println(maxPathSum(r))
}

func TestLevel(t *testing.T) {
	r := makeTree()

	res := levelOrder(r)

	for _, v := range res {
		fmt.Println(v)
	}

	fmt.Println(res)
}

func TestSubTree(t *testing.T) {
	r := makeTree()
	ar := makeAnotherTree()
	println(isSubtree(r, ar))
}

func TestBuildTree(t *testing.T) {
	// preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
	r := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	printTree(r)
}

func TestValidBST(t *testing.T) {
	r := makeBST()
	println(isValidBST(r))
}
