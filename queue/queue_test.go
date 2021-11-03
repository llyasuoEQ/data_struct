package queue

import "testing"

func TestQueue(t *testing.T) {
	q := NewQueueLinear()
	_ = q.EnQueue("111")
	_ = q.EnQueue("222")
	_ = q.EnQueue("333")
	e, _ := q.DeQueue()
	q.Print(e)
}
