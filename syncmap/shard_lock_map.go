package main

// 分片加锁

import (
	"sync"
)

type Player struct {
	LandUpdateNum             int32 // 更新计数 ，系统启动后 从0自增，每20次 会执行一次保存逻辑
	FactoryParallelUpgradeNum int32 // 吃物品 加经验 执行次数 %5 保存
	BuildActNum               int32 //  生产器点击累计次数

	// 下面这些暂定 基础信息模块 mod_basic_info
	UserId  int64
	Setting uint64 //客户端控制的
}

var ShardCount int64 = 32

// ShardLockMaps A "thread" safe map of type string:Anything.
// To avoid lock bottlenecks this map is dived to several (ShardCount) map shards.
type ShardLockMaps struct {
	shards []*SingleShardMap
}

// SingleShardMap A "thread" safe string to anything map.
type SingleShardMap struct {
	items map[int64]*Player
	sync.RWMutex
}

// createShardLockMaps Creates a new concurrent map.
func createShardLockMaps() ShardLockMaps {
	slm := ShardLockMaps{
		shards: make([]*SingleShardMap, ShardCount),
	}
	for i := 0; i < int(ShardCount); i++ {
		slm.shards[i] = &SingleShardMap{items: make(map[int64]*Player)}
	}
	return slm
}

// NewShardLockMaps Creates a new ShardLockMaps.
func NewShardLockMaps() ShardLockMaps {
	return createShardLockMaps()
}

// GetShard returns shard under given key
func (slm ShardLockMaps) GetShard(key int64) *SingleShardMap {
	return slm.shards[key%ShardCount]
}

// Count returns the number of elements within the map.
func (slm ShardLockMaps) Count() int {
	count := 0
	for i := int64(0); i < ShardCount; i++ {
		shard := slm.shards[i]
		shard.RLock()
		count += len(shard.items)
		shard.RUnlock()
	}
	return count
}

// Get retrieves an element from map under given key.
func (slm ShardLockMaps) Get(key int64) (*Player, bool) {
	if key <= 0 {
		return nil, false
	}
	shard := slm.GetShard(key)
	shard.RLock()
	val, ok := shard.items[key]
	shard.RUnlock()
	return val, ok
}

// Set Sets the given value under the specified key.
func (slm ShardLockMaps) Set(userData *Player) {
	if userData == nil {
		return
	}
	shard := slm.GetShard(userData.UserId)
	shard.Lock()
	shard.items[userData.UserId] = userData
	shard.Unlock()
}

// Remove removes an element from the map.
func (slm ShardLockMaps) Remove(key int64) {
	if key <= 0 {
		return
	}
	shard := slm.GetShard(key)
	shard.Lock()
	delete(shard.items, key)
	shard.Unlock()
}

// IsEmpty checks if map is empty.
func (slm ShardLockMaps) IsEmpty() bool {
	return slm.Count() == 0
}

// Keys returns all keys as []int64
func (slm ShardLockMaps) Keys() []int64 {
	count := slm.Count()
	ch := make(chan int64, count)
	go func() {
		wg := sync.WaitGroup{}
		wg.Add(int(ShardCount))
		for _, shard := range slm.shards {
			go func(shard *SingleShardMap) {
				shard.RLock()
				for key := range shard.items {
					ch <- key
				}
				shard.RUnlock()
				wg.Done()
			}(shard)
		}
		wg.Wait()
		close(ch)
	}()

	keys := make([]int64, 0, count)
	for k := range ch {
		keys = append(keys, k)
	}
	return keys
}

// IterCb Iterator callback,called for every key,value found in maps.
// RLock is held for all calls for a given shard
// therefore callback sess consistent view of a shard,
// but not across the shards
type IterCb func(key int64, v *Player)

// IterCb Callback based iterator, cheapest way to read
// all elements in a map.
func (slm ShardLockMaps) IterCb(fn IterCb) {
	for idx := range slm.shards {
		shard := slm.shards[idx]
		shard.RLock()
		for key, value := range shard.items {
			fn(key, value)
		}
		shard.RUnlock()
	}
}
