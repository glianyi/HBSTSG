package blind

// [1,8,6,2,5,4,8,3,7]
// 输出：49
func maxAreaView(height []int) int {
	maxArea, l, r := 0, 0, len(height)-1
	for l < r {
		area := minInt(height[l], height[r]) * (r - l)
		if area > maxArea {
			maxArea = area
		}
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}
	return maxArea
}
