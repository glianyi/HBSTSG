package blind

import "fmt"

func heapSort(nums []int) {
	h := heap{
		list: nums,
	}

	h.heapify()
	res := []int{}

	for len(h.list) > 0 {
		res = append(res, h.pop())
	}

	fmt.Println(res)
}
