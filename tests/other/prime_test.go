package other

import (
	"github.com/stretchr/testify/assert"
	"github.com/yb19890724/leetcode-go/src/other"
	"testing"
)

// @test prime
func TestPrime(t *testing.T)  {

	assert.Equal(t,other.Prime(11),true," prime 11 judge error ")
	assert.Equal(t,other.Prime(47),true," prime 47 judge error ")

}
