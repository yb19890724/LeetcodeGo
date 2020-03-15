package calculation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


// @test add two number value
func TestAddTwoNumbers(t *testing.T) {
	
	l1 := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 2,
				Next: nil,
			},
		},
	}
	
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
				
				},
			},
		},
	}
	
	l3 := &ListNode{
		Val: 8,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 7,
				Next: &ListNode{
				
				},
			},
		},
	}
	
	
	
	assert.Equal(t, l3 ,addTwoNumbers(l1,l2), " add wwo numbers value not equal ")
	
	
}