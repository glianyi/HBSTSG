package blind

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i := range nums {
		if _, exist := m[target-nums[i]]; exist {
			return []int{m[target-nums[i]], i}
		}
		m[nums[i]] = i
	}

	return []int{}
}
