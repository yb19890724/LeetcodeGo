package other

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// @test prime
func TestPrime(t *testing.T) {

	assert.Equal(t, Prime(11), true, " prime 11 judge error ")

	assert.Equal(t, Prime(47), true, " prime 47 judge error ")

}
