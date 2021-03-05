package sort

import (
	//"fmt"
	"github.com/yb19890724/leetcode-go/src/structure"
)

// 堆排序
// 理解：选择排序的一种优化

// 执行流程：
// 1.对数据堆化 O(n)
// 重复执行以下操作，直到堆的元素数量为1为止
// 1）.交换顶堆元素与尾元素   O(1)
// 2）.堆元素数量减1         O(1)
// 3) .对0位置进行1次siftDown操作  O(logn)

// 例：
//    0  1  2  3  4  5
// [ 50 21 80 43 38 14]
//             80
// 		   43┌────┴────┐50
//         ┌─┴─┐	 ┌─┘
//        21	38	 14


// @todo 优化空间 如果用堆来查找最小值 O(n*logn)
func HeapSort(data []int) []int {
	
	h := structure.MinHeap{
		Data:make([]int,len(data)),
	}
	
	h.HeapIfy(data)
	
	var res = make([]int,len(data))
	
	s:=h.Size
	for i :=0; i < s; i++ {
		t := h.Remove()
		res[i] =t
	}
	
	return res
}