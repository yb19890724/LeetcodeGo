package structure

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func createStack() ArrayStack {
	arrayStack := ArrayStack{
		10,
		[10]int{},
		0,
	}
	return arrayStack
}

// @test stack push
func TestStackPush(t *testing.T) {

	arrayStack := createStack()

	for i := 1; i <= arrayStack.MaxSize; i++ {
		assert.Equal(t, arrayStack.Push(i), true, " stack pop push ")
	}

}

// @test stack pop
func TestStackPop(t *testing.T) {

	arrayStack := createStack()

	var testStack [10]int

	for i := 0; i < arrayStack.MaxSize; i++ {
		testStack[i] = i + 1
		arrayStack.Push(i + 1)
	}

	for j := arrayStack.Top; j <= 1; j-- {
		assert.Equal(t, arrayStack.Pop(), testStack[arrayStack.Top-j], " stack pop error  ")
	}

}
