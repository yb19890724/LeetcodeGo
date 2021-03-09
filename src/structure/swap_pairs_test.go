package structure

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// @test Swap Pairs
func TestSwapPairs(t *testing.T) {

	data := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	l := SwapPairs(&data)

	for _, v := range []int{2, 1, 3, 4} {

		assert.Equal(t, v, v, " node value error")

		if l.Next == nil {
			break
		}
		l = l.Next
	}

}
