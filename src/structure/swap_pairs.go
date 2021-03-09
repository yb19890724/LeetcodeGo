package structure

// ListNode define

// 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

// Given 1->2->3->4, you should return the list as 2->1->4->3.
//输入：head = [1,2,3,4]
//输出：[2,1,4,3]

type ListNode struct {
	Val  int
	Next *ListNode
}

func SwapPairs(head *ListNode) *ListNode {
	// 思路：将链表翻转转化为一个子问题，然后通过递归方式依次解决
	// 先翻转两个，然后将后面的节点继续这样翻转，然后将这些翻转后的节点连接起来
	return helper(head)
}

// 递归过程
// 1 2 3 4
// 第一步 先取出3继续往下交换
// 第二部 2 和 1进行交换
// 第三部 把1的下个节点赋值成3和4交换完的节点

func helper(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 保存下一阶段的头指针
	nextHead := head.Next.Next

	// 翻转当前阶段指针
	next := head.Next
	next.Next = head

	head.Next = helper(nextHead)

	return next
}
