package blind

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

// 输入：[1,8,6,2,5,4,8,3,7]
// 49
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxV := min(height[left], height[right]) * (right - left)
	for left < right {
		if height[left] < height[right] {
			left += 1
		} else {
			right -= 1
		}
		area := min(height[left], height[right]) * (right - left)
		if area > maxV {
			maxV = area
		}
	}
	return maxV
}
