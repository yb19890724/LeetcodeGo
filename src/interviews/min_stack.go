package interviews

// 问：实现一个栈，要求实现Push（出栈）、Pop（入栈）、Min（返回最小值的操作）的时间复杂度为O(1)~

type Node struct {
	Data int
	Min  int
}

type Stack interface {
	Push()
	Pop()
	Min()
}

type MinStack struct {
	MaxSize int
	Top     int
	Array   [10]Node
}

// 入栈
func (stack *MinStack) Push(item int) bool {

	// 满栈
	if stack.Top == stack.MaxSize {

		return false

	}

	node := Node{item, item}

	if 0 == stack.Top { // 第一次入栈

		stack.Array[stack.Top] = node

	} else {

		stack.Array[stack.Top] = node

		min := stack.Min()

		if min < item {
			stack.Array[stack.Top].Min = min
		}

	}

	stack.Top++

	return true
}

// 出栈
func (stack *MinStack) Pop() int {

	if stack.Top == 0 {

		panic(" stack is nil ")

	}

	stack.Top--

	item := stack.Array[stack.Top]

	return item.Min
}

// 获取最小元素
func (stack *MinStack) Min() int {

	if stack.Top == stack.MaxSize {

		panic(" stack is nil ")

	}

	item := stack.Array[stack.Top-1].Min

	return item
}
