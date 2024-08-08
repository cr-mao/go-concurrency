/**
* @Author: maozhongyu
* @Desc:
* @Date: 2024/8/6
**/
package main

import (
	"bytes"
	"sync"
)

var buffers = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func GetBuffer() *bytes.Buffer {
	return buffers.Get().(*bytes.Buffer)
}

// bad
func PutBuffer(buf *bytes.Buffer) {
	buf.Reset()
	buffers.Put(buf)
}
func PutBuffer2(buf *bytes.Buffer) {
	// 大于 64kb 则不放回
	if buf.Cap() > 64<<10 {
		return
	}

	buf.Reset()
	buffers.Put(buf)
}
