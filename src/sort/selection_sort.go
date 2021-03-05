package sort

// 选择排序

// 执行顺序：
// 1.从序列中找出最小的元素，如果当前元素不是最小的元素，与当前元素与最小元素位置交换
// 2.下一次忽略上次的元素继续找剩下元素最小的

// 分析:选择排序的交换次数要远远少于冒泡排序，平均性能优于冒泡排序
// 最好，最坏，平均时间复杂度 O(n2) ,空间复杂度O(1)，属于稳定排序


// @todo 优化空间 如果用堆来查找最小值 O(n*logn)
func SelectionSort(data []int) []int {
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
