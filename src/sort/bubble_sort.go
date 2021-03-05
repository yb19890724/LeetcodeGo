package sort

// 冒泡排序
// 时间复杂度：最坏 平均 O(n²)
// 属于稳定排序

func BubbleSort(data []int) []int {

	for i := 0; i <= len(data); i++ {

		for j := i + 1; j < len(data); j++ {

			if data[j] < data[i] {

				data[i], data[j] = data[j], data[i]

			}
		}
	}

	return data
}

// @todo 冒泡排序优化版本
// 优化点：
// 如果序列已经局部有序，可以记录最后一次交换的位置减少比较次数
// 如果已经有序 最好 O(n)

func BubbleSort2(data []int) []int {
	
	for i := len(data) - 1; i > 0; i-- {
		
		index := 0
		
		for j := 1; j <= i ; j++ {
			
			if data[j] < data[j-1] {
				
				data[j-1], data[j] = data[j], data[j-1]
				
				index = j
			}
		}
		// 记录最后有序的交换位置
		i = index
	}
	
	return data
}