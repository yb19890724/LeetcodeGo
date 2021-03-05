package query

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// @test length of substring
func TestLengthOfSubString(t *testing.T) {
	
	assert.Equal(t, LengthOfSubString("aaabbabcd"), 4, " max check for errors ")
	
	assert.Equal(t, LengthOfSubString(""), 0, " max check for errors ")
}
