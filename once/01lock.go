package main

import (
	"net"
	"sync"
	"time"
)

var conMu sync.Mutex
var conn net.Conn

func getConn() net.Conn {
	if conn != nil {
		return conn
	}
	conMu.Lock()
	if conn != nil {
		return conn
	}
	defer conMu.Unlock()
	conn, _ = net.DialTimeout("TCP", ":80", 10*time.Second)
	return conn
}

func main() {
	conn := getConn()
	if conn == nil {
		panic("conn is nil")
	}
}
