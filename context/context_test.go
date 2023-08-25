package context_test

import (
	"context"
	"testing"
	"time"
)

func Test_ContextValue(t *testing.T) {
	var ctx = context.Background()
	ctx = context.WithValue(ctx, "key1", "0001")
	ctx = context.WithValue(ctx, "key2", "0002")
	ctx = context.WithValue(ctx, "key3", "0003")
	ctx = context.WithValue(ctx, "key4", "0004")
	t.Log(ctx.Value("key1"))
	// 查找key ->key4:0004->key3:0003->key2:0002->key1:0001
}

type key struct{}

// 推荐用这种
func Test_ContextValue2(t *testing.T) {
	var ctx = context.Background()
	ctx = context.WithValue(ctx, key{}, "0001")
	t.Log(ctx.Value(key{}))
}

func Test_ContextCancel(t *testing.T) {
	var ctx = context.Background()
	ctx, cancel := context.WithCancel(ctx)
	t.Log(ctx.Deadline()) // 0001-01-01 00:00:00 +0000 UTC false
	go func() {
		time.Sleep(time.Second * 5)
		cancel()
	}()
	<-ctx.Done()
	t.Logf("context close error: %v", ctx.Err())
}

func Test_ContextDeadline(t *testing.T) {
	var ctx = context.Background()
	t.Log(time.Now().Unix())
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second*5))
	deadTime, ok := ctx.Deadline() // ok 表示是否设置过期时间
	t.Log(deadTime, ok)
	defer cancel()
	<-ctx.Done()
	t.Log(time.Now().Unix()) //于上面相差5秒
	t.Logf("context close error: %v", ctx.Err())
}

//本质和 WithDeadline是一样的
func Test_ContextWithTimeout(t *testing.T) {
	var ctx = context.Background()
	t.Log(time.Now().Unix())
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	<-ctx.Done()
	t.Log(time.Now().Unix()) //于上面相差5秒
	t.Logf("context close error: %v", ctx.Err())
}
