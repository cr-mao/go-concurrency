/**
* @Author: maozhongyu
* @Desc:
* @Date: 2024/7/12
**/
package main

import "sync"

type Mutex[T any] struct {
	mu    sync.Mutex
	value T
}

func NewMutex[T any](value T) *Mutex[T] {
	return &Mutex[T]{
		value: value,
	}
}

func (m *Mutex[T]) Load() T {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.value
}

func (m *Mutex[T]) Store(value T) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.value = value
}

func (m *Mutex[T]) Lock(f func(v *T)) {
	m.mu.Lock()
	defer m.mu.Unlock()
	//f(&m.value)
	value := m.value
	f(&value)
	m.value = value
}
