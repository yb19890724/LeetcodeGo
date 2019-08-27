package structure

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 初始化公用队列结构
func initCircularQueue() *CircularArrayQueue {
	circularQueue := &CircularArrayQueue{
		5,
		[5]int{},
		0,
		0,
	}

	circularQueue.Enqueue(1)
	circularQueue.Enqueue(2)
	circularQueue.Enqueue(3)
	circularQueue.Enqueue(4)

	return circularQueue
}

// @test  queue enqueue 测试入列
func TestCircularEnQueue(t *testing.T) {

	CircularArrayQueue := initCircularQueue()

	assert.Equal(t, CircularArrayQueue.All(), [5]int{1, 2, 3, 4, 0}, " queue enqueue error ")

}

// @test  queue dequeue 测试出队
func TestCircularDequeue(t *testing.T) {

	CircularArrayQueue := initCircularQueue()

	testData := []int{1, 2, 3, 4, 0}

	for i := 0; i < CircularArrayQueue.MaxSize-1; i++ {

		assert.Equal(t, CircularArrayQueue.Dequeue(), testData[i], " queue dequeue error ")

	}

}

// @test queue max size 测试队列长度限制
func TestCircularMaxSize(t *testing.T) {

	CircularArrayQueue := initQueue()

	assert.Equal(t, CircularArrayQueue.Enqueue(6), false, " queue max size overstep ")

}
