package types

// Stack defines the struct of a FILO stack
type Stack struct{
	data []interface{}
}

// Push pushes an item into the stack
func (s *Stack) Push(x interface{}) {
	s.data = append(s.data, x)
}

// Pop pops an item out from the top of the stack
func (s *Stack) Pop() (x interface{}) {
	if len(s.data) == 0 {
		return
	}

	x = s.data[len(s.data)-1]
	s.data = s.data[0 : len(s.data)-1]
	return
}

// Peek returns the item at the top of the stack
// and doesn't make any change the stack itself
func (s *Stack) Peek() (x interface{}) {
	if len(s.data) == 0 {
		return
	}

	x = s.data[len(s.data)-1]
	return
}
