package queue

import (
	"errors"
	"fmt"
	"sync"
)

type Ele interface{}

type QueueLinear struct {
	eles []Ele
	mu   sync.Mutex
}

// NewQueueLinear
func NewQueueLinear() *QueueLinear {
	return &QueueLinear{
		eles: []Ele{},
		mu:   sync.Mutex{},
	}
}

// Print 打印
func (q *QueueLinear) Print(e Ele) error {
	if q == nil {
		return errors.New("QueueLinear is nil")
	}
	fmt.Println(q.eles)
	return nil
}

// EnQueue 进队列
func (q *QueueLinear) EnQueue(e Ele) error {
	if q == nil {
		return errors.New("QueueLinear is nil")
	}
	q.mu.Lock()
	q.eles = append(q.eles, e)
	q.mu.Unlock()
	return nil
}

// DeQueue 出队列
func (q *QueueLinear) DeQueue() (Ele, error) {
	if q == nil {
		return nil, errors.New("QueueLinear is nil")
	}
	if q.IsEmpty() {
		return nil, errors.New("queue have not ele")
	}
	e := q.eles[0]
	q.mu.Lock()
	q.eles = q.eles[1:]
	defer q.mu.Unlock()
	return e, nil
}

// IsEmpty 是否为空
func (q *QueueLinear) IsEmpty() bool {
	if q == nil || len(q.eles) == 0 {
		return false
	}
	return true
}
