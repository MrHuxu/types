package types

// Queue defines the struct of a FIFO queue
type Queue struct {
	data []interface{}
}

// Size returns the size of the queue
func (q *Queue) Size() int {
	return len(q.data)
}

// Enqueue puts an item to the tail of the queue
func (q *Queue) Enqueue(x interface{}) {
	q.data = append(q.data, x)
}

// Dequeue gets an item from the head of the queue
func (q *Queue) Dequeue() (x interface{}) {
	if len(q.data) == 0 {
		return
	}

	x = q.data[0]
	q.data = q.data[1:len(q.data)]
	return
}

// Peek returns the item at the head of the queue
// and doesn't make any change the queue itself
func (q *Queue) Peek() (x interface{}) {
	if len(q.data) == 0 {
		return
	}

	x = q.data[0]
	return
}
