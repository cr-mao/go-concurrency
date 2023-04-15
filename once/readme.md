## sync.Once

lock 实现也行，就是一直用锁，性能开销过大。 其实第二次可以不用互斥锁了。 判断值即可


自定义once，保证func 执行成功，没有err 。 04advance_once.go


