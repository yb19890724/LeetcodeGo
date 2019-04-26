package interviews

import (
	"github.com/stretchr/testify/assert"
	"github.com/yb19890724/leetcode-go/src/interviews"
	"testing"
)

// @test min stack push
func TestMinStackPush(t *testing.T) {

	var elem int = 1

	var item [10]interviews.Node

	item[0] = interviews.Node{elem, elem}

	minStack := interviews.MinStack{10, 0, [10]interviews.Node{}}

	judge := minStack.Push(elem)

	assert.Equal(t, judge, true, " min stack push error ")

	assert.Equal(t, minStack.Array, item, " min stack array error ")

}

// @test min stack pop
func TestMinStackPop(t *testing.T) {


	minStack := interviews.MinStack{10, 0, [10]interviews.Node{}}

	minStack.Push(1)

	assert.Equal(t, minStack.Pop(), 1, " min stack pop error ")

}

// @test min stack min
func TestMinStackMin(t *testing.T) {


	minStack := interviews.MinStack{10, 0, [10]interviews.Node{}}

	minStack.Push(1)
	minStack.Push(2)
	minStack.Push(3)

	assert.Equal(t, minStack.Min(), 1, " min stack get min error ")

	minStack.Pop()

	assert.Equal(t, minStack.Min(), 1, " min stack get min error ")

}
