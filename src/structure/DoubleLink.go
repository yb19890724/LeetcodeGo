package structure

import "fmt"

type DoubleNode struct {
	PrevNode *DoubleNode
	NextNode *DoubleNode
	Data     int
}

type DoubleLink struct {
	Size     uint
	HeadNode *DoubleNode
	TailNode *DoubleNode
}

func (d *DoubleLink) length() uint {

	return d.Size
}

// 头部增加节点
func (d *DoubleLink) Add(data int) {
	newNode := &DoubleNode{Data: data, PrevNode: nil, NextNode: nil}

	// 空链表
	if 0 == d.length() {
		d.HeadNode = newNode
		d.TailNode = newNode

		return
	}
	// 头部节点
	d.HeadNode.PrevNode = newNode
	newNode.NextNode = d.HeadNode
	d.HeadNode = newNode

	d.Size += 1
}

// 尾部增加节点
func (d *DoubleLink) Append(data int) {

	newNode := &DoubleNode{Data: data, PrevNode: nil, NextNode: nil}

	if 0 == d.length() {
		d.HeadNode = newNode
		d.TailNode = newNode
		d.Size += 1
		return
	}

	tail := d.TailNode
	tail.NextNode = newNode
	newNode.PrevNode = tail
	d.TailNode = newNode

	d.Size += 1
}

func (d *DoubleLink) Reverse() {

	curr := d.TailNode

	for nil != curr {

		fmt.Println(curr.Data)

		if nil != curr.PrevNode {
			curr = curr.PrevNode
		} else {
			break
		}

	}

}

func (d *DoubleLink) Print() {

	curr := d.HeadNode

	for nil != curr {

		fmt.Println(curr.Data)

		if nil != curr.NextNode {
			curr = curr.NextNode
		} else {
			break
		}
	}

}
