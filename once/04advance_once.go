package main

import (
	"sync"
	"sync/atomic"
)

type Once struct {
	done uint32
	m    sync.Mutex
}

func (o *Once) Do(fn func() error) error {
	if atomic.LoadUint32(&o.done) == 1 {
		return nil
	}
	return o.doSlow(fn)
}

func (o *Once) doSlow(fn func() error) error {
	o.m.Lock()
	defer o.m.Unlock()
	var err error
	if o.done == 0 {
		err = fn()
		if err == nil { // 只有初始化成功才打标注
			atomic.StoreUint32(&o.done, 1)
		}
	}
	return err
}
