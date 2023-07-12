package blind

// https://leetcode.cn/problems/contains-duplicate/solution/chao-xiang-xi-kuai-lai-miao-dong-ru-he-p-sf6e/
func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for _, v := range nums {
		if _, exists := m[v]; exists {
			return true
		}
		m[v] = true
	}

	return false
}
