package types

// Queue defines the struct of a FIFO queue
type Queue struct {
	data []interface{}
}

// Enqueue puts an item to the tail of the queue
func (s *Queue) Enqueue(x interface{}) {
	s.data = append(s.data, x)
}

// Dequeue gets an item from the head of the queue
func (s *Queue) Dequeue() (x interface{}) {
	if len(s.data) == 0 {
		return
	}

	x = s.data[0]
	s.data = s.data[1:len(s.data)]
	return
}

// Peek returns the item at the head of the queue
// and doesn't make any change the queue itself
func (s *Queue) Peek() (x interface{}) {
	if len(s.data) == 0 {
		return
	}

	x = s.data[0]
	return
}
