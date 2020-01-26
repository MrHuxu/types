package types

// Stack defines the struct of a FILO stack
type Stack []interface{}

// Push pushes an item into the stack
func (s Stack) Push(x interface{}) {
	s = append(s, x)
}

// Pop pops an item out from the top of the stack
func (s Stack) Pop() (x interface{}) {
	if len(s) == 0 {
		return
	}

	x = s[len(s)-1]
	s = s[0 : len(s)-1]
	return
}

// Peek returns the item at the top of the stack
// and doesn't make any change the stack itself
func (s Stack) Peek() (x interface{}) {
	if len(s) == 0 {
		return
	}

	x = s[len(s)-1]
	return
}
