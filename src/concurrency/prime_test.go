package concurrency

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// @test Concurrency prime
func TestConcurrencyPrime(t *testing.T) {

	prime := Sieve()

	data := []int{2, 3}

	for i := 0; i < len(data); i++ {

		assert.Equal(t, <-prime, data[i], " prime judge error")

	}

}
