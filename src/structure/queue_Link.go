package structure

import "fmt"

type QueueElem struct {
	Data int
	Next *QueueElem
}

type QueueLink struct {
	Data int
	List *QueueElem
	Size int
}

func (q *QueueLink) Enqueue(data int) {

	elem := &QueueElem{Data: data, Next: nil}

	if 0 == q.Size {

		q.List = elem
		q.Size++

		return
	}

	curr := q.List

	for nil != curr {

		if nil != curr.Next {
			curr = curr.Next
		}

	}

	curr = elem

	q.List = curr
	q.Size++
}

func (q *QueueLink) Dequeue() {

	tmp := q.List
	q.List = q.List.Next
	fmt.Println(tmp)
}
