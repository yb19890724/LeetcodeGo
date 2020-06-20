package other

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 递归版本
func TestFibonacci(t *testing.T) {

	data := [10]int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}

	for i := 1; i <= 10; i++ {

		assert.Equal(t, Fibonacci(i), data[i-1], " fibonacci error")

	}
}


// 非递归版本
func TestFibonacci2(t *testing.T) {
	
	data := [10]int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	
	for i := 1; i <= 10; i++ {
		
		fmt.Println(Fibonacci(i))
		assert.Equal(t, Fibonacci2(i), data[i-1], " fibonacci error")
		
	}
}

