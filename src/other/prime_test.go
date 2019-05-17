package other

import (
	"github.com/stretchr/testify/assert"
	"github.com/yb19890724/leetcode-go/src/goruntine"
	"testing"
)

// @test prime
func TestPrime(t *testing.T) {

	assert.Equal(t, Prime(11), true, " prime 11 judge error ")

	assert.Equal(t, Prime(47), true, " prime 47 judge error ")

}

// @test goruntime prime
func TestGoRuntimePrime(t *testing.T) {

	prime := goruntine.Sieve()

	data := []int{2, 3}

	for i := 0; i < len(data); i++ {

		assert.Equal(t, <-prime, data[i], " prime judge error")

	}

}
