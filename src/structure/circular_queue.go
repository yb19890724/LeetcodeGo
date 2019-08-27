package structure

/**
 * circular queue 循环队列
 * @todo 利用数组存储队列，把首尾相连形成一个环
 *  循环数组会浪费一个空间
 * -------------------------------------------------------------
 * [思路说明]
 * Head    指向头部
 * Tail    指向尾部
 * Array   存储数据
 * MaxSize 队列最大长度
 * -------------------------------------------------------------
 */

type CircularQueue interface {
	Enqueue() // 入列
	Dequeue() // 出列
	Show()    // 遍历队列
	All()
}

// 数组实现：顺序队列
type CircularArrayQueue struct {
	MaxSize int
	Array   [5]int
	Head    int // 头指针
	Tail    int // 尾指针
}

// 入列
func (queue *CircularArrayQueue) Enqueue(item int) (boolean bool) {

	// 判断队列是否满了
	if (queue.Tail+1)%queue.MaxSize == queue.Head {
		return false
	}

	queue.Array[queue.Tail] = item

	queue.Tail = (queue.Tail + 1) % queue.MaxSize

	return true
}

// 出列
func (queue *CircularArrayQueue) Dequeue() (item int) {

	// 判断队列是否是空
	if queue.Head == queue.Tail {

		panic(" queue is nil ")

	}

	index := queue.Head

	queue.Head = (queue.Head + 1) % queue.MaxSize

	return queue.Array[index]
}

// 获取队列所有值
func (queue *CircularArrayQueue) All() [5]int {

	return queue.Array

}
