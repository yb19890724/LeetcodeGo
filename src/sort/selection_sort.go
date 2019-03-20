package sort

// 选择排序
func SelectionSort(data []int) []int{
	times := 0
	for i := 0; i < len(data)-1; i++ {
		min := i
		for j := i; j < len(data); j++ {
			times++
			if data[min] > data[j] {
				min = j
			}
		}
		if min != i {
			tmp := data[i]
			data[i] = data[min]
			data[min] = tmp
		}
	}
	return data
}

