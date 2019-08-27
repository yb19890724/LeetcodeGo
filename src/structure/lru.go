package structure

import (
	"container/list"
	"fmt"
)

// @todo :LRU缓存淘汰算法
// LRU是最近最少使用策略的缩写，是根据数据的历史访问记录来进行淘汰数据，其核心思想是“如果数据最近被访问过，那么将来被访问的几率也更高”.
// 使用has map ➕ 双链表实现.
// 使用List保存数据,Map来做快速访问即可.

// lru数据结构
type LruCache struct {
	capacity int                      // cache长度
	Items    map[string]*list.Element //元素
	Link     *list.List
}

// 元素结构
type Item struct {
	Key   string
	Value interface{}
}

// 创建cache数据结构
func NewLruCache(capacity int) *LruCache {
	return &LruCache{
		capacity: capacity,
		Items:    make(map[string]*list.Element),
		Link:     list.New(),
	}
}

// @todo:获取缓存数据  通过list更新cache key,在hap map中获取数据
// 1.判断元素是否在list存在，不存在返回空接口<nil>
// 2.如果存在，将元素移动到链表的最前面
// 3.从hap map中获取数据，并且返回结果

func (l *LruCache) Get(key string) (value interface{}, exists bool) {
	if item, exists := l.Items[key]; exists {
		l.Link.MoveToFront(item)
		value = item.Value.(*Item).Value
	}
	return
}

// @todo:设置元素 list和has map
// 1.如果元素存在，把元素移动到list头结点，并且更新hap map中的元素值
// 2.如果不存在，在list最前面插入元素，并且写入到hap map中
// 3.如果list的长度超过了设置长度，则删除list最后一个元素
func (l *LruCache) Set(key string, value interface{}) {
	if item, exists := l.Items[key]; exists {
		// 移到头结点并更新值
		l.Link.MoveToFront(item)
		item.Value.(*Item).Value = value
		return
	}
	item := l.Link.PushFront(&Item{key, value})
	l.Items[key] = item
	if l.Link.Len() > l.capacity {
		l.DeleteOldest()
	}
}

// @todo:删除list历史数据
// 1.返回链表最后一个元素，如果list为空，返回nil（如果list没有满的情况返回的是nil）
// 2.如果存在值，就从list中删除元素
func (l *LruCache) DeleteOldest() {
	item := l.Link.Back() // returns the last element of list or nil if the list is empty.
	if item != nil {
		l.removeItem(item)
	}
}

// @todo:删除元素
func (l *LruCache) delete(key string) {
	if item, exists := l.Items[key]; exists {
		l.removeItem(item)
	}
}

// @todo: 移除list中记录，和hap map存储的元素值
func (l *LruCache) removeItem(el *list.Element) {

	l.Link.Remove(el)

	item := el.Value.(*Item)

	delete(l.Items, item.Key)
}

// @todo: 打印第一个元素，如果不是nil，继续往下便利
func (l *LruCache) print() {
	tmp := l.Link.Front()
	for tmp != nil {
		fmt.Printf("%v -> ", tmp.Value)
		tmp = tmp.Next()
	}
	fmt.Println()
}
