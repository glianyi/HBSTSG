package blind

func climbStairs(n int) int {
	if n == 0 || n == 1 || n == 2 {
		return n
	}

	return climbStairs(n-1) + climbStairs(n-2)
}

func climbStairsIter(n int) int {
	if n == 0 || n == 1 || n == 2 {
		return n
	}

	res := make([]int, n)
	res[0], res[1] = 1, 2
	for i := 2; i < n; i++ {
		res[i] = res[i-1] + res[i-2]
	}

	return res[n-1]
}

func climbStairsIter2(n int) int {
	if n == 0 || n == 1 || n == 2 {
		return n
	}

	a, b := 1, 2

	for i := 0; i < n-2; i++ {
		a, b = b, a+b
	}

	return b
}
