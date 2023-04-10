package main

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	cap     int                   // 容量
	size    int                   // 当前缓存大小
	cache   map[int]*list.Element // 哈希表存储键值对
	lruList *list.List            // 双向链表键值对
}

type Pair struct {
	key   int
	value int
}

// NewLRUCache 创建一个 LRUCache 对象
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		cap:     capacity,
		cache:   make(map[int]*list.Element),
		lruList: list.New(),
	}
}

// Get 获取一个键值对的 value，如果存在则将其移到双向链表头部
func (c *LRUCache) Get(key int) int {
	if elem, ok := c.cache[key]; ok {
		c.lruList.MoveToFront(elem) // 如果查询的元素存在，将把它放到链表的第一位
		return elem.Value.(*Pair).value
	}
	return -1
}

// Put 将k，v放入哈希表和双向链表
func (c *LRUCache) Put(key, value int) {
	if elem, ok := c.cache[key]; ok {
		elem.Value.(*Pair).value = value
		c.lruList.MoveToFront(elem) // 如果放入的元素重复，或是换value，就是使用了这个key，将这个元素放到第一个位置
	} else { // 如果哈希表没有这个key
		if c.size == c.cap { // 首先判断缓存满了没有
			delete(c.cache, c.lruList.Back().Value.(*Pair).key) // 从哈希表删除最后一个元素
			c.lruList.Remove(c.lruList.Back())                  // 再从双向链表删除最后一个元素
			c.size--                                            // 缓存大小减1
		}
		c.cache[key] = c.lruList.PushFront(&Pair{key, value}) // 将新的k, v放入双向链表第一个位，并且将其放入哈希表
		c.size++
	}
}

func main() {
	lruCache := NewLRUCache(2)
	lruCache.Put(1, 1)
	lruCache.Put(2, 2)
	fmt.Println(lruCache.Get(1)) // 1
	lruCache.Put(3, 3)
	fmt.Println(lruCache.Get(2)) // -1
	lruCache.Put(4, 4)
	fmt.Println(lruCache.Get(1)) // -1
	fmt.Println(lruCache.Get(3)) // 3
	fmt.Println(lruCache.Get(4)) // 4
}
