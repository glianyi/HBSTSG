package blind

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	fmt.Println(nums)

	res := make([][]int, 0)

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			continue
		}
		left := i + 1
		right := len(nums) - 1

		sum := nums[i] + nums[left] + nums[right]
		for left < right {
			if sum == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				println(nums[i], nums[left], nums[right])
				for ; left < right && nums[left] == nums[left+1]; left++ {
				}
				for ; left < right && nums[right] == nums[right-1]; right-- {
				}
				left++
				right--
			}

			if sum < 0 {
				left++
			}

			if sum > 0 {
				right--
			}
		}

	}
	return res
}
