package main

import (
	"fmt"
	"sync"
	"time"
)

type SliceQueue struct {
	data []interface{}
	mu   sync.Mutex
}

func NewQueue(n int) *SliceQueue {
	return &SliceQueue{
		data: make([]interface{}, 0, n),
	}
}

// 入队 把值放在队尾
func (q *SliceQueue) Enqueue(v interface{}) {
	q.mu.Lock()
	q.data = append(q.data, v)
	q.mu.Unlock()
}

// 出队，移除队头并返回
func (q *SliceQueue) Dequeue() interface{} {
	q.mu.Lock()

	// 如果队列为空，直接返回nil
	if len(q.data) == 0 {
		q.mu.Unlock()
		return nil
	}
	v := q.data[0]
	// 移除队头
	q.data = q.data[1:]
	q.mu.Unlock()
	return v
}

func main() {
	queue := NewQueue(100)
	go func() {
		for i := 0; i < 100; i++ {
			queue.Enqueue(i)
		}
	}()
	for {
		v := queue.Dequeue()
		fmt.Println("出队的值是：", v)
		time.Sleep(50 * time.Millisecond)
	}
}
