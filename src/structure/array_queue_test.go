package structure

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 初始化公用队列结构
func initQueue() *ArrayQueue {
	arrayQueue := &ArrayQueue{
		5,
		[5]int{},
		0,
		0,
	}
	
	arrayQueue.Enqueue(1)
	arrayQueue.Enqueue(2)
	arrayQueue.Enqueue(3)
	arrayQueue.Enqueue(4)
	arrayQueue.Enqueue(5)
	
	return arrayQueue
}

// @test  queue enqueue 测试入列
func TestEnQueue(t *testing.T) {
	
	arrayQueue := initQueue()
	
	assert.Equal(t, arrayQueue.All(), [5]int{1, 2, 3, 4, 5}, " queue enqueue error ")
	
}

// @test  queue dequeue 测试出队
func TestDequeue(t *testing.T) {
	
	arrayQueue := initQueue()
	
	testData := []int{1, 2, 3, 4, 5}
	
	for i := 0; i < arrayQueue.MaxSize; i++ {
		assert.Equal(t, arrayQueue.Dequeue(), testData[i], " queue dequeue error ")
	}
	
}

// @test queue max size 测试队列长度限制
func TestQueueMaxSize(t *testing.T) {
	
	arrayQueue := initQueue()
	
	assert.Equal(t, arrayQueue.Enqueue(6), false, " queue max size overstep ")
	
}
