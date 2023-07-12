package blind

func findMin(nums []int) int {
	left, right, mid := 0, len(nums)-1, -1
	for left < right {
		mid = (left + right) / 2
		if nums[mid] > nums[right] {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return nums[mid]
}
