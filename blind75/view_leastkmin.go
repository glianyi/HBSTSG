package blind

// [2,5,3,7,5,4]
// [2,3,5,5,4,7]
// [2,3,5,4,5,7]
// [2,3,4,5,5,7]
func bubbleSort(arr []int) []int {
	swap := false
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swap = true
			}
		}
		if !swap {
			break
		}
	}

	return arr
}

func insertSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i; j > 1; j-- {
			if arr[j] < arr[j-1] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
	return arr
}

func selectSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		maxI := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[maxI] {
				maxI = j
			}
		}
		arr[i], arr[maxI] = arr[maxI], arr[i]
	}

	return arr
}

func quickSort(arr []int, l, r int) []int {
	if l >= r {
		return arr
	}
	partition := func() int {
		pivot := r
		i := l
		for j := l; j < r; j++ {
			if arr[j] < arr[pivot] {
				arr[j], arr[i] = arr[i], arr[j]
				i++
			}
		}

		arr[pivot], arr[i] = arr[i], arr[pivot]
		return i
	}

	i := partition()
	quickSort(arr, l, i-1)
	quickSort(arr, i+1, r)
	return arr
}

func getLeastNumbers(arr []int, k int) []int {
	swap := false
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swap = true
			}
		}
		if !swap || i == len(arr)-k {
			break
		}
	}

	return arr[:k]
}
