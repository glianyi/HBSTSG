package blind

// 输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
// 6
func maxSubArray(nums []int) int {
	maxValue := nums[0]
	accumlate := nums[0]
	for right := 1; right < len(nums); right++ {
		accumlate += nums[right]
		if accumlate > maxValue {
			maxValue = accumlate
		}

		if nums[right] > accumlate {
			maxValue = nums[right]
			accumlate = maxValue
		}

	}
	return maxValue
}
