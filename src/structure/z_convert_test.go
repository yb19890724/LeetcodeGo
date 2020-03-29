package structure

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// @test z convert
func TestZConvert(t *testing.T) {
	
	assert.Equal(t, ZConvert("LEETCODEISHIRING", 3), "LCIRETOESIIGEDHN", " string not equal errors ")
	
	assert.Equal(t, ZConvert("LEETCODEISHIRING", 4), "LDREOEIIECIHNTSG", " string not equal errors ")
}
