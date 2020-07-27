package sort

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
	times := 0
	for i := 0; i < len(data)-1; i++ {
		// 最小的元素索引
		min := i
		for j := i; j < len(data); j++ {
			times++
			if data[min] > data[j] {
				min = j
			}
		}
		
		// 如果当前索引不是最小的索引，那就把最小的索引与现有所以值交换
		if min != i {
			tmp := data[i]
			data[i] = data[min]
			data[min] = tmp
		}
	}
	return data
}