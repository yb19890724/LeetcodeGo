package structure

import "fmt"

// 循环链表
// 关键点是 让尾指针指向头指针

type CycleLinkNode struct {
	Data int
	Next *CycleLinkNode
}

// 约瑟夫环问题

// 重点是构造一个循环列表，循环报数的时候，加一个临时变量存储的是前一个列表的节点，为了删除指向
func WhoIsKing(n int) {

	// 循环链表
	cycleLink := initCycLink(n)

	///当前报的数是几
	count := 0

	// 用来记录报数的节点
	curr := new(CycleLinkNode)

	// 退出条件不是最后一个元素就需要一直报数
	for cycleLink != cycleLink.Next {
		// 模拟报数
		count++
		// 判断当前报数是不是3
		if 3 == count {
			// 如果是3 把当前节点删除调用，打印当前节点，重置报数初始化，让当前节点后移

			// 删除节点
			curr.Next = cycleLink.Next
			// 打印删除的节点
			fmt.Print(cycleLink.Data,",")
			// 每次循环以后挪动节点
			cycleLink = cycleLink.Next
			count = 0
			continue
		}

		// 不是3 当前节点后移
		curr = cycleLink
		// 循环遍历往后挪
		cycleLink = cycleLink.Next
	}

	// 最后一个节点
	fmt.Print(curr.Data)
}

func ShowCycleLinkLink(n int) {
	curr := initCycLink(n)

	i := 1

	for nil != curr {

		fmt.Println(curr.Data)

		if nil != curr.Next {
			curr = curr.Next
		}

		// 退出
		if i == n {
			break
		}

		i++
	}

}

func initCycLink(n int) *CycleLinkNode {

	head := new(CycleLinkNode)
	prev := new(CycleLinkNode)

	for i := 1; i <= n; i++ {

		newNode := &CycleLinkNode{
			Data: i,
			Next: nil,
		}

		// 第一次构建
		if 1 == i {

			// 头结点
			head = newNode
			// 前一节节点就是头节点
			prev = head

			continue
		}

		// 当前节点的下一个节点是新节点
		prev.Next = newNode
		// 当前节点变成新的节点
		prev = newNode

		// 构建完成了把最后一个节点指向头节点
		if i == n {
			prev.Next = head
		}

	}

	return head
}

// 快指针：是慢指针的两倍
// 中间值问题   1，2，3，4，5，6，7 （获得4）
// 是否有环
// 链表入口

// @todo  单链表中间数问题
// 便利过程
// 第一步
// slow 和 fast 都指向 aa
// aa bb cc dd ee ff gg

// 第二部步
// slow 指向 bb
// fast 指向 cc
// aa bb cc dd ee ff gg

// 遍历到最后  slow正好指向中位值
// slow 指向 dd
// fast 指向 gg
// aa bb cc dd ee ff gg

func (l *Link) LinkMiddle() {
	var slow, fast *LinkNode

	slow = l.HeadNode
	fast = l.HeadNode

	for nil != fast && nil != fast.Next {

		slow = slow.Next
		fast = fast.Next.Next

	}

}

// @todo 单链表是否有环

// 无环：是指前一个元素指向后一个元素
// aa -> bb -> cc -> dd -> ee -> ff -> gg

// 有环：某一个节点是从后往前指的  关键点是后面节点指向前面的节点
// aa -> bb -> cc -> dd -> ee -> ff -> (gg -> dd)
// https://www.bilibili.com/video/BV1M7411f7aU?p=22

// 关键点定义快慢指针，如果说快指针追上了慢指针那么则有环存在

func (l *Link) LinkIsCycle() bool {

	var slow, fast *LinkNode

	slow = l.HeadNode
	fast = l.HeadNode

	for nil != fast && nil != fast.Next {

		slow = slow.Next
		fast = fast.Next.Next

		if slow.Data == fast.Data {
			return true
		}

	}

	return false
}

// @todo 找到有环的入口位置
// 当快慢指针相遇时，我们可以判断链表有环，这时设定一个新指针指向链表的起点（重点），且步长与慢指针一样，则(慢指针)与(新指针)相遇的地方就是环的入口。

func (l *Link) searchEnter() {
	var temp, fast, slow *LinkNode

	fast = l.HeadNode
	slow = l.HeadNode

	for nil != fast && nil != fast.Next {

		fast = fast.Next.Next
		slow = slow.Next

		if fast.Data == slow.Data {
			temp = l.HeadNode
		}

		if nil != temp {
			temp = temp.Next

			if temp.Data == slow.Data {
				return
			}
		}
	}
}
