## sync.Pool

 sync.Pool 临时对象池 （对象缓存）

- 尝试从私有对象获取
- 私有对象不存在，尝试从当前processor的共享池获取
- 如果当前processor共享池也是空的，那么从其他processor的共享池获取
- 如果所有子池都是空的，那么就用用户指定的new函数产生一个新的对象返回
- 私有对象 协程安全，共享池协程不安全


- Sync.Pool对象的放回
- 如果私有对象不存在则保存私有对象
- 如果私有对象存在，放入当前Processor子池的共享池



**sync.Pool 受gc 影响** 见 gc_pool.go



Get方法实现： 

```go
func (p *Pool) Get() any {
	if race.Enabled {
		race.Disable()
	}
	//把当前groutine 固定在当前P上
	l, pid := p.pin()
	x := l.private  //优先从 private 字段取，快速
	l.private = nil
	if x == nil { // private 不存在
		// Try to pop the head of the local shard. We prefer
		// the head over the tail for temporal locality of
		// reuse.   
		//从当前local shared 弹出一个，注意从head 读取并移除
		x, _ = l.shared.popHead()
		if x == nil {
			//如果没有 取其他P上偷一个
			x = p.getSlow(pid)
		}
	}
	runtime_procUnpin()
	if race.Enabled {
		race.Enable()
		if x != nil {
			race.Acquire(poolRaceAddr(x))
		}
	}
	
	//都没有拿到，尝试New 函数去生产一个新的
	if x == nil && p.New != nil {
		x = p.New()
	}
	return x
}

```


Put 实现

```go
// Put adds x to the pool.
func (p *Pool) Put(x any) {
	if x == nil { // nil 直接丢弃
		return
	}
	if race.Enabled {
		if fastrandn(4) == 0 {
			// Randomly drop x on floor.
			return
		}
		race.ReleaseMerge(poolRaceAddr(x))
		race.Disable()
	}
	l, _ := p.pin()
	if l.private == nil { //如果本地private 为空， 则直接设置这个指即可
		l.private = x
		x = nil
	}
	if x != nil { // 否则加入到本地队列
		l.shared.pushHead(x)
	}
	runtime_procUnpin()
	if race.Enabled {
		race.Enable()
	}
}
```


### sync.Pool的使用陷阱

看一段 bufpool.go  这里其实有内存浪费的电。

如果要放回元素的 cap 很大， 那么所占用的空间依然很大。

解决办法： 将元素放回时，增加检测逻辑，如果要放回的元素超过一定大小的 buffer，就直接丢弃，不再放回池子。





### tcp连接池

fatih/pool   其实是稳定的。基于 channel 实现




### memcached Client 连接池
还可以用 切片维护连接池
https://github.com/bradfitz/gomemcache/blob/master/memcache/memcache.go

freelist[:len(freelist)-1]

put的使用 append(freelist, cn)


### net/rpc 中的 Request/Response 对象池

使用链表使用