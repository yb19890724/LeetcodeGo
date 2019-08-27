package sort

// 冒泡排序
// 时间复杂度：O(n²)

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
