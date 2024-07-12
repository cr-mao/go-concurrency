/**
* @Author: maozhongyu
* @Desc: 并发泛型 map
* @Date: 2024/7/12
**/
package main

import (
	"fmt"
	"sync"
)

type Map[k comparable, v any] struct {
	mu    sync.Mutex
	value map[k]v
}

// 初始化 map
func NewMap[k comparable, v any](size ...int) *Map[k, v] {
	if len(size) > 0 {
		return &Map[k, v]{
			value: make(map[k]v, size[0]),
		}
	}
	return &Map[k, v]{
		value: make(map[k]v),
	}
}

func (m *Map[k, v]) Get(key k) (v, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	val, ok := m.value[key]
	// 没有， val 是零值
	return val, ok
}

func (m *Map[k, v]) Set(key k, val v) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.value[key] = val
}

func main() {
	var m1 = NewMap[string, int]()
	m1.Set("mao", 35)
	fmt.Println(m1.Get("mao"))
}
