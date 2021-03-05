package query

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// @test longest palindrome
func TestLongestPalindrome(t *testing.T) {
	
	assert.Equal(t, LongestPalindrome("babad"), "aba", " max check for errors ")
	
	assert.Equal(t, LongestPalindrome("cbbd"), "bb", " max check for errors ")
	
	assert.Equal(t, LongestPalindrome("ac"), "c", " max check for errors ")
	
	assert.Equal(t, LongestPalindrome("ccc"), "ccc", " max check for errors ")
}
