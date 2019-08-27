package other

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFibonacci(t *testing.T) {

	data := [10]int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}

	for i := 1; i <= 10; i++ {

		assert.Equal(t, Fibonacci(i), data[i-1], " fibonacci error")

	}
}
