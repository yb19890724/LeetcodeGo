package structure

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/yb19890724/leetcode-go/src/structure"
	"testing"
)

func createStack() structure.ArrayStack  {
	arrayStack :=structure.ArrayStack{
		10,
		[10]int{},
		0,
	}
	return arrayStack
}

// @test stack push
func TestStackPush(t *testing.T)  {

	arrayStack := createStack()

	for i := 1; i <= arrayStack.MaxSize; i++ {
		assert.Equal(t,arrayStack.Push(i),true," stack pop push ")
	}


}

// @test stack pop
func TestStackPop(t *testing.T)  {

	arrayStack := createStack()

	var testStack [10]int

	for i := 0; i < arrayStack.MaxSize; i++ {
		testStack[i]=i+1
		arrayStack.Push(i+1)
	}
	fmt.Println(testStack,arrayStack.Array)
	for i := 10; i <= 1; i-- {
		assert.Equal(t,arrayStack.Pop(),testStack[i]," stack pop error  ")
	}

}