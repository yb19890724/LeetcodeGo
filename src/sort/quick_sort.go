package sort

// 快速排序
// 时间复杂度：
// 最好 O(n*logn) 平均 O(n*logn) 最坏 O(n2)

// 流程分析：
// 1.以第一个数为中间轴进行递归切割 比轴大和比轴小的
// 2.切到1个元素进行合病操作左面的肯定比右面的小


func QuickSort(data []int) []int {

	if len(data) <= 1 {

		return data

	}

	left, right := make([]int, 0, len(data)), make([]int, 0, len(data))

	for i := 1; i < len(data); i++ {

		if data[i] <= data[0] {

			left = append(left, data[i])

		}

		if data[i] > data[0] {

			right = append(right, data[i])

		}
	}

	left, right = QuickSort(left), QuickSort(right)

	left = append(left, data[0])

	return append(left, right...)
}
