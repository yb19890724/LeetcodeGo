package structure

import (
	"container/list"
	"fmt"
)

// lru数据结构
type LruCache struct {
	capacity int
	items    map[string]*list.Element
	link     *list.List
}

type Item struct {
	key   string
	value interface{}
}

// 创建cache数据结构
func NewLruCache(capacity int) *LruCache {
	return &LruCache{
		capacity: capacity,
		items:    make(map[string]*list.Element),
		link:     list.New(),
	}
}

// 获取缓存数据
func (l *LruCache) Get(key string) (value interface{}, exists bool) {
	if item, exists := l.items[key]; exists {
		l.link.MoveToFront(item) // moves element e to the front of list l. If e is not an element of l, the list is not modified. The element must not be nil.
		value = item.Value.(*Item).value
	}
	return
}

func (l *LruCache) Set(key string, value interface{}) {
	if item, exists := l.items[key]; exists {
		// 移到头结点并更新值
		l.link.MoveToFront(item)
		item.Value.(*Item).value = value
		return
	}
	item := l.link.PushFront(&Item{key, value}) // inserts a new element at the front of list
	l.items[key] = item
	if l.link.Len() > l.capacity {
		l.DeleteOldest()
	}
}

func (l *LruCache) DeleteOldest() {
	item := l.link.Back() // returns the last element of list or nil if the list is empty.
	if item != nil {
		l.removeItem(item)
	}
}

func (l *LruCache) delete(key string) {
	if item, exists := l.items[key]; exists {
		l.removeItem(item)
	}
}

func (l *LruCache) removeItem(el *list.Element) {
	l.link.Remove(el) // removes el from list if el is an element of the list. It returns the element value el.Value. The element must not be nil.

	item := el.Value.(*Item)
	delete(l.items, item.key)
}

func (l *LruCache) print() {
	tmp := l.link.Front() // returns the first element of list or nil if the list is empty.
	for tmp != nil {
		fmt.Printf("%v -> ", tmp.Value)
		tmp = tmp.Next() // returns the next list element
	}
	fmt.Println()
}

func main() {
	lru := NewLruCache(5)
	lru.Get("s")
	lru.print()
	lru.Set("o", 90)
	lru.Set("n", 200)
	lru.Set("y", "fujifilm")
	lru.Get("n")
	lru.print() // &{n 200} -> &{y fujifilm} -> &{o 90} ->
	lru.Set("x", "xt3")
	lru.Set("gfx", "50s")
	lru.Set("sony", "a7r")
	lru.Get("gfx")
	lru.print() // &{gfx 50s} -> &{sony a7r} -> &{x xt3} -> &{n 200} -> &{y fujifilm} ->
}