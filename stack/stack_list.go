package stack

import (
	"errors"
	"fmt"
	"sync"
)

type Ele interface{} // 栈元素

// StackLinear stackLinear 利用线性的结构实现的栈
type StackLinear struct {
	eles []Ele
	mu   sync.Mutex
}

func NewStackLinear() *StackLinear {
	return &StackLinear{
		eles: []Ele{},
		mu:   sync.Mutex{},
	}
}

// Print 打印
func (s *StackLinear) Print(e Ele) {
	fmt.Println(e)
}

// Push 压栈
func (s *StackLinear) Push(e Ele) error {
	if s == nil {
		return errors.New("StackLinter is nil")
	}
	s.mu.Lock()
	s.eles = append(s.eles, e)
	s.mu.Unlock()
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
	l := s.Len()
	e := s.eles[l-1]
	s.mu.Lock()
	s.eles = s.eles[0 : l-1]
	s.mu.Unlock()
	return e, nil
}

// IsEmpty 是否为空
func (s *StackLinear) IsEmpty() bool {
	if s == nil || len(s.eles) == 0 {
		return false
	}
	return true
}

// Len 栈的长度
func (s *StackLinear) Len() int {
	if s == nil || len(s.eles) == 0 {
		return 0
	}
	return len(s.eles)
}
