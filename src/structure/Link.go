package structure

type LinkNode struct {
	Data int
	Next *LinkNode
}

type Link struct {
	HeadNode *LinkNode // 头节点
}

func (l *Link) isEmpty() bool {

	if nil == l.HeadNode {
		return true
	}

	return false
}

func (l *Link) Add(d int) *LinkNode {

	tmp := l.HeadNode

	l.HeadNode = &LinkNode{d, tmp}

	return l.HeadNode
}

// 尾部添加
func (l *Link) Append(d int) {

	n := &LinkNode{Data: d, Next: nil}

	if l.isEmpty() {
		l.HeadNode = n
		return
	}

	curr := l.HeadNode

	for {

		if nil != curr.Next {
			curr = curr.Next
		} else {
			break
		}
	}

	curr.Next = n

}

// 移除指定index位置元素
func (l *Link) Remove(index int) {

	n := 0

	if l.isEmpty() {
		return
	}

	if index == 0 {
		tmp := l.HeadNode.Next
		l.HeadNode = tmp
		return
	}

	curr := l.HeadNode

	if nil == curr {
		return
	}

	for nil != curr.Next {

		if nil != curr.Next {
			curr = curr.Next
		}

		if n == index {
			break
		}

		n++
	}

	curr.Next = curr.Next.Next

}
