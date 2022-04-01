package main

import "fmt"

//func longestPalindrome2(s string) string {
//	res := ""
//	for i := 0; i < len(s); i++ {
//		res = maxPalindrome(s, i, i, res)
//		res = maxPalindrome(s, i, i+1, res)
//	}
//	return res
//}
//
//func maxPalindrome(s string, i, j int, res string) string {
//	sub := ""
//	for i >= 0 && j < len(s) && s[i] == s[j] {
//		sub = s[i : j+1]
//		i--
//		j++
//	}
//
//	if len(res) < len(sub) {
//		return sub
//	}
//	return res
//}
//
//
//
//
//
//
//
//
//

func longestPalindrome(s string) string {

	res := ""
	for i := 0; i < len(s); i++ {
		res = maxPalindrome(i, i, s, res)
		res = maxPalindrome(i, i+1, s, res)
	}

	return res
}

func maxPalindrome(i int, j int, s string, res string) string {
	sub := ""
	for i >= 0 && j < len(s) && s[i] == s[j] {
		sub = s[i : j+1]
		i--
		j++
	}

	if len(sub) > len(res) {
		return sub
	}
	return res
}

func main() {

	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	root := removeNthFromEnd(head, 2)

	fmt.Println(root)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	nidNode := &ListNode{Val: 0,Next: head}

	curr,pre := nidNode,nidNode

	for curr.Next != nil {

		if curr.Val == curr.Next.Val {

			pre.Next = curr.Next.Next

			curr = pre

			continue
		}

		pre = curr

		curr = curr.Next
	}

	return nidNode
}
