package blind

func insert(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0)
	inserted := false

	for _, iv := range intervals {
		if inserted {
			res = append(res, iv)
			continue
		}

		if newInterval[0] > iv[1] {
			res = append(res, iv)
			continue
		}

		if newInterval[1] < iv[0] {
			res = append(res, newInterval)
			res = append(res, iv)
			inserted = true
			continue
		}

		newInterval[0] = minInt(iv[0], newInterval[0])
		newInterval[1] = maxInt(iv[1], newInterval[1])
	}

	return res
}
