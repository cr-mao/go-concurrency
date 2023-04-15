package main

import (
	"context"
	"fmt"
)

func main() {
	var ctx = context.Background()
	ctx = context.WithValue(ctx, "key1", "0001")
	ctx = context.WithValue(ctx, "key2", "0002")
	ctx = context.WithValue(ctx, "key3", "0003")
	ctx = context.WithValue(ctx, "key4", "0004")

	fmt.Println(ctx.Value("key1"))

	// 查找key ->key4:0004->key3:0003->key2:0002->key1:0001
}
