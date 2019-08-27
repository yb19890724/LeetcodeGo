package structure

/**
 * Array queue
 * 利用数组实现顺序队列
 * @todo 注意数组空间利用，在入栈时重置元素位置，让空出来的空间可以循环使用.
 * -------------------------------------------------------------
 * [思路说明]
 * Head    指向头部
 * Tail    指向尾部
 * Array   存储数据
 * MaxSize 队列最大长度
 * -------------------------------------------------------------
 */

// 队列接口
type Queue interface {
	Enqueue() // 入列
	Dequeue() // 出列
	Show()    // 遍历队列
	All()
}

// 数组实现：顺序队列
type ArrayQueue struct {
	MaxSize int
	Array   [5]int
	Head    int // 头指针
	Tail    int // 尾指针
}

// 入列
func (queue *ArrayQueue) Enqueue(item int) (boolean bool) {

	// 入列前先判断队列是否已满
	if queue.Tail == queue.MaxSize {

		// Tail ==n && Head==0，表示整个队列都占满了
		if queue.Head == 0 {
			return false
		}

		// 移动数据
		// 插入数据的时候每一次，每次把数据都重置到第一位0开始存放数据
		for i := queue.Head; i < queue.Tail; i++ {
			queue.Array[i-queue.Head] = queue.Array[i]
		}

		// 搬移完之后重新更新Head和Tail
		queue.Tail -= queue.Head // 指向前一位 做后面的出处
		queue.Head = 0
	}

	queue.Array[queue.Tail] = item // 入列

	queue.Tail++ // 指针后移

	return true
}

// 出列
func (queue *ArrayQueue) Dequeue() (item int) {

	// 判断队列是否是空
	if queue.Head == queue.Tail {

		panic(" queue is nil ")

	}

	index := queue.Head

	queue.Head++

	return queue.Array[index]
}

// 获取队列所有值
func (queue *ArrayQueue) All() [5]int {

	return queue.Array

}
