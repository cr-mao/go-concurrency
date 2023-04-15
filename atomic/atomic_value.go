package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type Config struct {
	NodeName string
	Addr     string
	Count    int
}

func loadNewConfig() Config {
	return Config{
		NodeName: "北京",
		Addr:     "10.0.0.1",
		Count:    rand.Intn(2),
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var config atomic.Value
	var cond = sync.NewCond(&sync.Mutex{})
	config.Store(loadNewConfig())
	// 设置新的config
	go func() {
		for {
			time.Sleep(time.Second)
			newConfig := loadNewConfig()
			c := config.Load().(Config)
			//只有配置变le 才通知打印 ，所以看到的都是不同的count 值
			if newConfig != c {
				config.Store(newConfig)
				cond.Broadcast()
			}
		}
	}()
	go func() {
		for {
			cond.L.Lock()
			cond.Wait()
			c := config.Load().(Config)
			fmt.Printf("new config %+v \n", c)
			cond.L.Unlock()
		}
	}()
	select {}
}
