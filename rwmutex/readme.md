
## sync.RWMutex
大量并发读， 并发写少的场景，考虑用RWMutex 代替Mutex 

写锁相关：  Lock() 、 TryLock() 、UnLock()

读锁相关: RLock() 、TryRLock()、RUnlock()


注意点:
- 不可复制 
- 重入导致死锁

    1,2之前哪怕是读锁， 1，2中间有写锁进来， 1这个读锁就永远释放不掉了。 
 ```go
mu.RLock(); // 1
{
mu.RLock()  // 2
fmt.Println("hello")
mu.RUnlock()
}
RUnlock
```

- 释放未加锁的RWMutex


发现死锁， 只要把标准库sync包 ，import的时候 替换 github.com/sasha-s/go-deadlock
