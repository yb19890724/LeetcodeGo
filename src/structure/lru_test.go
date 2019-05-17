package structure

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Describer interface {
	Describe()
}
type Person struct {
	name string
	age  int
}

func (p Person) Describe() {
	fmt.Printf("%s is %d years old", p.name, p.age)
}

func findType(i interface{}) {
	switch v := i.(type) {
	case Describer:
		v.Describe()
	default:
		fmt.Printf("unknown type\n")
	}
}

// @var test lru cache
// 存储值时，a最开始在末尾，按顺序呢存入，超过后被淘汰
// 但是a在中间被访问所以，b在最后面被淘汰调：满足最近最少使用策略
func TestLruCache(t *testing.T) {

	lru := NewLruCache(5)

	lru.Set("a", 1)
	lru.Set("b", 2)
	lru.Set("c", 3)
	lru.Get("a")
	lru.Set("e", 4)
	lru.Set("f", 5)
	lru.Set("g", 6)

	data := []string{"g", "f", "e", "a", "c"}

	tmp := []int{6, 5, 4, 1, 3}

	for key, value := range data {

		if item, exists := lru.Items[value]; exists {

			assert.Equal(t,tmp[key],item.Value.(*Item).Value)

		}

	}
}
