package main

import "fmt"

func main() {
	var m = make(map[string]int)
	m["a"] = 0
	fmt.Printf("a=%d,b=%d\n", m["a"], m["b"]) // 0,0

	av, aexists := m["a"]
	bv, bexists := m["b"]
	fmt.Printf("a=%d,a existd=%t, b=%d ,b existd=%t", av, aexists, bv, bexists)
}
