package blind

import "fmt"

// nums = [1,2,3,4]
// [1,1,2,6]
// [24,12,4,1]
// 输出: [24,12,8,6]

func productExceptSelf(nums []int) []int {
	left := []int{1}
	right := []int{1}

	for i := 0; i < len(nums)-1; i++ {
		left = append(left, left[len(left)-1]*nums[i])
		right = append([]int{right[0] * nums[len(nums)-i-1]}, right...)
	}

	fmt.Println(left, right)
	for i := 0; i < len(nums); i++ {
		left[i] = left[i] * right[i]
	}
	return left
}
