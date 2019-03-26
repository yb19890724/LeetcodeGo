package structure

// 栈是一种“操作受限”的线性表，只允许在一端插入和删除数据。
// 后进先出
// 顺序栈

type Stack interface {
	Push() // 入栈
	Pop()  // 出栈
}

type ArrayStack struct {
	MaxSize int   // 最大长度
	Array   [10]int // 数组
	Count   int   // 元素数量
}

// 入栈
func (stack *ArrayStack) Push(item int) bool {

	// 判断栈是否满了
	if stack.Count == stack.MaxSize {

		return false

	}

	stack.Array[stack.Count] = item

	item++

	return true
}

// 出栈
func (stack *ArrayStack) Pop() int {

	if stack.Count == 0 {
		panic(" stack is null ")
	}

	item := stack.Array[stack.Count-1]

	stack.Count--

	return item
}
