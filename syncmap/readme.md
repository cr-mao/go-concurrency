## sycn.Map


sync.RWMutex 扩展map 


分片加锁 map . https://github.com/orcaman/concurrent-map 



sync.Map 提供了 9 个方法，我们可以把它们归为三类

- 读操作
  - Load(key any) (value any,ok bool) 读取一个键对应的值
  - Range(f func(key,value any)) 遍历 map。
- 写操作
  - Store(key,value any) : 存储或者更新一个键.
  - Delete(key any) 删除一个键
  - Swap(key,value any) （previous any,loaded bool) : 替换一个键 ，并把以前的结果返回。如果这个键不存在，loaded 返回 false，但新值还是设置成功
- 读/写操作
  - CompareAndDelete(key,old any) (deleted bool) 如果锁提供的值和旧值相等，则删除这个键
  - CompareAndSwap(key,old,new any) (swapped bool) ： cas 操作，如果所提供的值和旧值相等，则设置新值。
  - LoadAndDelete(key any) (value any,loaded bool) 返回并删除一个键，如果这个键不存在，则 loaded 返回 false 
  - LoadOrStore(key,value any) (actual any,loaded bool) ：如果这个键不存在，则设置新值,loaded返回 false，并返回新值。 如果这个键存在，则返回这个键对应的值,loaded=true。



lock-free map: https://github.com/alphadose/haxmap 