package blind

import "fmt"

type heap struct {
	list []int
}

//          100(0)
//         /      \
//      40(1)     50(2)
//     /  \      /      \
//   10(3) 15(4) 30(5) 40(6)

func (h *heap) heapify() { // nlog(n)
	for i := 0; i < len(h.list); i++ {
		for j := len(h.list); j > i; {
			p := h.parent(j)
			maxIndex := h.whichMax(p)
			h.list[maxIndex], h.list[p] = h.list[p], h.list[maxIndex]
			j = p
		}
		// fmt.Println(h.list)
	}
}

func (h *heap) push(e int) {
	h.list = append(h.list, e)
	for i := len(h.list) - 1; i > 0; {
		p := h.parent(i)
		if h.list[p] > h.list[i] {
			return
		}
		h.list[i], h.list[p] = h.list[p], h.list[i]
		i = p
	}
	return
}

func (h *heap) pop() int {
	if len(h.list) == 0 {
		return -1
	}
	res := h.list[0]
	h.list[0] = h.list[len(h.list)-1]
	h.list = h.list[:len(h.list)-1]
	for i := 0; i < len(h.list); {
		maxIndex := h.whichMax(i)
		if maxIndex == i { // 说明不能找到比这个值更小的值了
			break
		}
		h.list[maxIndex], h.list[i] = h.list[i], h.list[maxIndex]
		i = maxIndex
	}
	return res
}

func (h *heap) peek() int {
	return h.list[0]
}

func (h *heap) whichMax(p int) int {
	l := h.left(p)
	r := h.right(p)
	maxIndex := p
	if l != -1 && h.list[l] > h.list[maxIndex] {
		maxIndex = l
	}

	if r != -1 && h.list[r] > h.list[maxIndex] {
		maxIndex = r
	}

	return maxIndex
}

func (h *heap) parent(i int) int {
	if i == 0 {
		return -1
	}
	return (i - 1) / 2
}

func (h *heap) left(i int) int {
	index := 2*i + 1
	if index >= len(h.list) {
		return -1
	}
	return index
}

func (h *heap) right(i int) int {
	index := 2*i + 2
	if index >= len(h.list) {
		return -1
	}
	return 2*i + 2
}

func (h *heap) printHeap() {
	fmt.Println(h.list)
}
