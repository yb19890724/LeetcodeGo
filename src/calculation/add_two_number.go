package calculation

// 给出两个 非空 的链表用来表示两个非负的整数。
// 其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
// 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

/*
示例
输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

*/

// 解决方案
// 我们首先从最低有效位也就是列表 l1 和 l2 的表头开始相加。由于每位数字都应当处于 0…9 的范围内，我们计算两个数字的和时可能会出现“溢出”。
// 例如，5 + 7 = 12。在这种情况下，我们会将当前位的数值设置为 2，并将进位 flag = 1 带入下一次迭代。进位 flag 必定是 0 或 1，这是因为两个数字相加（考虑到进位）可能出现的最大和为 9 + 9 + 1 = 19。


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// 两数相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	
	if l1 == nil && l2 == nil {
		return initListNode(0)
	}
	
	var value, flag int
	var res, curr *ListNode
	
	
	// flag 为了照顾 5+5 =10  的情况
	for  l1 != nil || l2 != nil || flag !=0 {
		
		sum := 0
		
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		
		// 溢出增加
		if flag != 0 {
			sum += flag
		}
		
		// 每个节点只能存储一位数字 余位
		value = sum % 10
		
		// (11~19)/10 = 1
		flag = sum / 10
		
		if res != nil {
			curr.Next = initListNode(value)
			curr = curr.Next
		} else {
			res = initListNode(value)
			curr = res
		}
		
	}
	
	return res
}


func initListNode(n int) *ListNode {
	node := &ListNode{}
	node.Val = n
	node.Next = nil
	return node
}