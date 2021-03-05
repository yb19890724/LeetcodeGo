package calculation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// @test add  number value
func TestAddNumbers(t *testing.T) {
	
	assert.Equal(t, AddNumber(10), 55, " add numbers value not equal ")
	
}

// @test add  number value
func TestAddNumbers2(t *testing.T) {
	
	assert.Equal(t, AddNumber2(10), 55, " add numbers value not equal ")
	
}