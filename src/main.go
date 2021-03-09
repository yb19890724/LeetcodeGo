package main

import (
	"github.com/yb19890724/leetcode-go/src/structure"
)

func main() {

	data := structure.ListNode{
		Val: 1,
		Next: &structure.ListNode{
			Val: 2,
			Next: &structure.ListNode{
				Val: 3,
				Next: nil,
			},
		},
	}
	t := structure.SwapPairs(&data)

	for {
		if t.Next == nil {
			break
		}
		t = t.Next
	}
}
