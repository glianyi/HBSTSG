package golang

func RangeVal() {
	nums := []int{1, 2, 3, 4}
	m := make(map[int]*int)

	for k, v := range nums {
		m[k] = &v
	}

	for k, v := range m {
		println(k, *v)
	}
}
