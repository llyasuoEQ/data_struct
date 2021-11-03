package stack

import (
	"errors"
	"fmt"
)

type Ele interface{} // 栈元素

// stackLinear 利用线性的结构实现的栈
type StackLinear struct {
	eles []Ele
}

func NewStackSlice() *StackLinear {
	return &StackLinear{
		eles: []Ele{},
	}
}

// Print 打印
func (s *StackLinear) Print() error {
	if s == nil {
		return errors.New("StackLinter is nil")
	}
	fmt.Println(s.eles)
	return nil
}

// Push 压栈
func (s *StackLinear) Push(e Ele) error {
	if s == nil {
		return errors.New("StackLinter is nil")
	}
	s.eles = append(s.eles, e)
	return nil
}

// Pop 出栈
func (s *StackLinear) Pop() (Ele, error) {
	if s == nil {
		return nil, errors.New("StackLinter is nil")
	}
	if !s.IsEmpty() {
		return nil, errors.New("StackLinter have not ele")
	}
	e := s.eles[0]
	s.eles = s.eles[1:]
	return e, nil
}

// IsEmpty 是否为空
func (s *StackLinear) IsEmpty() bool {
	if s == nil || len(s.eles) == 0 {
		return false
	}
	return true
}
