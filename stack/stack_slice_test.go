package stack

import "testing"

func TestStack(t *testing.T) {
	s := NewStackLinear()
	_ = s.Push("111")
	_ = s.Push("222")
	_ = s.Push("333")
	e, _ := s.Pop()
	s.Print(e)
}
