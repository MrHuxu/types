package types

// MinHeap defines the struct of a heap
// which has the minimal item at the top
type MinHeap struct {
	data []int
}

// Size returns the size of the heap
func (m *MinHeap) Size() int {
	return len(m.data)
}

// Push puts an item into the heap
// and keep the minimal item at the top of heap
func (m *MinHeap) Push(x int) {
	m.data = append(m.data, x)
	reorder(m.data, isSmaller)

	return
}

// Pop puts an item into the heap
// and keep the minimal item at the top of heap
func (m *MinHeap) Pop() (x int) {
	if len(m.data) == 0 {
		return
	}

	x = m.data[0]
	m.data = m.data[1:len(m.data)]
	reorder(m.data, isSmaller)

	return
}

// Peek returns the item at the top of the heap
// and doesn't make any change the stack itself
func (m *MinHeap) Peek() (x int) {
	if len(m.data) == 0 {
		return
	}

	x = m.data[0]
	return
}

// MaxHeap defines the struct of a heap
// which has the maximum item at the top
type MaxHeap struct {
	data []int
}

// Size returns the size of the heap
func (m *MaxHeap) Size() int {
	return len(m.data)
}

// Push puts an item into the heap
// and keep the minimal item at the top of heap
func (m *MaxHeap) Push(x int) {
	m.data = append(m.data, x)
	reorder(m.data, isLarger)

	return
}

// Pop puts an item into the heap
// and keep the minimal item at the top of heap
func (m *MaxHeap) Pop() (x int) {
	if len(m.data) == 0 {
		return
	}

	x = m.data[0]
	m.data = m.data[1:len(m.data)]
	reorder(m.data, isLarger)

	return
}

// Peek returns the item at the top of the heap
// and doesn't make any change the stack itself
func (m *MaxHeap) Peek() (x int) {
	if len(m.data) == 0 {
		return
	}

	x = m.data[0]
	return
}

func reorder(data []int, check func(int, int) bool) {
	for i := len(data) / 2; i >= 1; i-- {
		parent := i - 1
		if left := i*2 - 1; !check(data[parent], data[left]) {
			tmp := data[parent]
			data[parent] = data[left]
			data[left] = tmp
		}

		if right := i * 2; right < len(data) && !check(data[parent], data[right]) {
			tmp := data[parent]
			data[parent] = data[right]
			data[right] = tmp
		}
	}
}

func isSmaller(a, b int) bool {
	return a < b
}

func isLarger(a, b int) bool {
	return a > b
}
