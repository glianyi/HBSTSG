package blind

import (
	"fmt"
	"testing"
)

func makeTree() *TreeNode { // -10 20 9 7 15
	// 	-10
	// 	/ \
	// 9  20
	// 	 /  \
	// 	15   7

	// 	  -10
	// 	  / \
	//   20  9
	// 	/  \
	// 7   15
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
	println()
	rr := buildTreeView([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	printTree(rr)
	println()
}

func TestValidBST(t *testing.T) {
	r := makeBST()
	println(isValidBST(r))
}

func TestLengthOfLongestSubstring(t *testing.T) {
	println(lengthOfLongestSubstring("abcab"))
	println(lengthOfLongestSubstring("abcde"))
	println(lengthOfLongestSubstring("bbbbbb"))
	println(lengthOfLongestSubstring("pwwkew"))
	println(lengthOfLongestSubstring("abcdeaf"))
	println(lengthOfLongestSubstring("aabbb"))
}

// AABABAA
func TestCharReplace(t *testing.T) {
	println(characterReplacement("AABABAA", 1))
	println(characterReplacement("AABABAA", 2))
	println(characterReplacement("AABABAACCCCCC", 2))
	println(characterReplacement("BAAAAAAAAAAB", 2))
}

func TestMinWindow(t *testing.T) {
	println(minWindow("ADOBECODEBANC", "ABC"))
	println(minWindow("A", "A"))
	println(minWindow("A", "AA"))
}

func TestIsAnagram(t *testing.T) {
	println(isAnagram("nihao", "haoni"))
	println(isAnagram("nihao", "haowo"))
}

func TestIsValid(t *testing.T) {
	println(isValid("(){}[]"))
	println(isValid("({[]})"))
	println(isValid("((){}[]"))
}

func TestGroupAnagram(t *testing.T) {
	g := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagram(g))
	// 输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
}

//	     100(0)
//	    /      \
//	 40(1)     50(2)
//	/  \      /      \
//
// 10(3) 15(4) 30(5) 40(6)
func TestHeap(t *testing.T) {
	h := heap{
		list: []int{100, 400, 50, 10, 15, 30, 40, 600},
	}
	h.printHeap()
	h.heapify()
	h.printHeap()
	println(h.pop())
	h.printHeap()
	println(h.peek())
	h.printHeap()
	h.push(1000)
	h.printHeap()
}

func TestHeapSort(t *testing.T) {
	heapSort([]int{8, 7, 4, 6, 2, 1, 5})
}

func TestIsPalindrome(t *testing.T) {
	println(isPalindrome("nihaooahin"))
	println(isPalindrome("nih,ao oahin"))
	println(isPalindrome("nih,aoo ahin"))
	println(isPalindrome("A man, a plan, a canal Panama"))
}

func TestLongestPalindrome(t *testing.T) {
	println(longestPalindrome("sbacabc"))
}

func TestMap(t *testing.T) {
	m := map[string]string{}
	m["a"] = "a"

	for k, v := range m {
		println(k, v)
	}
}

func TestCountSubstrings(t *testing.T) {
	println(countSubstrings("sbacabc"))
}

func TestClimbStairs(t *testing.T) {
	println(climbStairs(1))
	println(climbStairsIter(1))
	println(climbStairsIter2(1))
	println(climbStairs(2))
	println(climbStairsIter(2))
	println(climbStairsIter2(2))
	println(climbStairs(3))
	println(climbStairsIter(3))
	println(climbStairsIter2(3))
	println(climbStairs(4))
	println(climbStairsIter(4))
	println(climbStairsIter2(4))
	println(climbStairs(5))
	println(climbStairsIter(5))
	println(climbStairsIter2(5))
}

func TestCoinchange(t *testing.T) {
	println(coinChange([]int{1, 2, 5}, 11))
	println(coinChange([]int{2}, 1))
	println(coinChange([]int{2, 5, 7}, 8))
	println(coinChange([]int{1}, 8))
}

// [10,9,2,5,3,7,101,18],4
func TestLongestlis(t *testing.T) {
	println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	println(lengthOfLISView([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}

func TestLongestCommonSubsequence(t *testing.T) {
	println(longestCommonSubsequence("abcbdab", "bdcaba"))
	println(longestCommonSubsequenceView("abcbdab", "bdcaba"))
}

func TestWordbreak(t *testing.T) {
	println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
	println(wordBreak("catdog", []string{"cats", "dog", "sand", "and", "cat"}))
}

func TestCombinationSum4(t *testing.T) {
	println(combinationSum4([]int{3}, 9))
	println(combinationSum4([]int{1, 2, 3}, 4))
}

func TestRob(t *testing.T) {
	println(rob([]int{1, 2, 3, 1}))
}

func TestCycleRob(t *testing.T) {
	println(cycleRob([]int{1, 2, 3, 1}))
	println(cycleRob([]int{2, 3, 2}))
	println(cycleRob([]int{1, 2, 3}))
	println(cycleRob([]int{1}))
	println(cycleRob([]int{1, 2}))
}

func TestUniquePaths(t *testing.T) {
	println(uniquePaths(7, 3))
	println(dp_uniquePath(7, 3))
}

func TestCanJump(t *testing.T) {
	println(canJump([]int{2, 3, 1, 1, 4}))
	println(canJump([]int{3, 2, 1, 0, 4}))
}

func TestTwoSum(t *testing.T) {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

func TestMaxProfit(t *testing.T) {
	println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	println(maxProfit([]int{7, 6, 4, 3, 1}))
}

func TestContainDup(t *testing.T) {
	println(containsDuplicate([]int{1, 3, 4, 2, 5, 1}))
}

func TestProductExceptSelf(t *testing.T) {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}

func TestMaxSubarray(t *testing.T) {
	println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	println(maxSubArrayDp([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	println(maxSubArray([]int{1}))
	println(maxSubArrayDp([]int{1}))
	println(maxSubArray([]int{5, 4, -1, 7, 8}))
	println(maxSubArrayDp([]int{5, 4, -1, 7, 8}))
}

// 输入: nums = [2,3,-2,4]
func TestMaxProduct(t *testing.T) {
	println(maxProduct([]int{2, 3, -2, 4}))
}

func TestFindmin(t *testing.T) {
	println(findMin([]int{4, 5, 6, 1, 2, 3}))
}

func TestMaxArea(t *testing.T) {
	println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	println(maxAreaView([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	println(maxArea([]int{1, 1}))
	println(maxAreaView([]int{1, 1}))
}

func TestThreeSum(t *testing.T) {
	fmt.Println(threeSum([]int{1, -1, -1, -4, 0, 0, 2}))
}

func TestInsert(t *testing.T) {
	// [[1,5],[6,9]]
	fmt.Println(insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))
	// 输出：[[1,2],[3,10],[12,16]]
	fmt.Println(insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}))
}

func TestMaxSum(t *testing.T) {
	// nums = [-2,1,-3,4,-1,2,1,-5,4] 连续子数组 [4,-1,2,1] 的和最大为6
	println(maxSum([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	println(maxSumDp([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

func TestMirrorTree(t *testing.T) {
	r := makeTree()
	printTree(r)
	println()
	i := invertTree(r)
	printTree(mirrorTree(r))
	println()
	// i := invertTreeIter(r)
	printTree(i)
	println()

	printTree(mirrorTreeIter(r))
	println()

	printTree(mirrorTreeL(r))
	println()
}

func TestNumWays(t *testing.T) {
	println(numWays(7))
	println(numWaysIter(7))
	println(numWaysIter(7))
}

func TestBubbleSort(t *testing.T) {
	fmt.Println(bubbleSort([]int{2, 5, 3, 7, 5, 4}))
	fmt.Println(getLeastNumbers([]int{2, 5, 3, 7, 4, 5}, 2))
	fmt.Println(insertSort([]int{2, 5, 3, 7, 5, 4}))
	fmt.Println(selectSort([]int{2, 5, 3, 7, 5, 4}))
	fmt.Println(quickSort([]int{2, 5, 3, 7, 5, 4}, 0, 5))
}
