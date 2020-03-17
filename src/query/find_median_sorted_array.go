package query

import "github.com/yb19890724/leetcode-go/src/sort"

// 给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
//
// 请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
//
// 你可以假设 nums1 和 nums2 不会同时为空。
//
// 示例 1:
//
// nums1 = [1, 3]
// nums2 = [2]
//
// 则中位数是 2.0
// 示例 2:
//
// nums1 = [1, 2]
// nums2 = [3, 4]
//
// 则中位数是 (2 + 3)/2 = 2.5

// 什么是两个数组的中位数?

// 1.奇数
// nums1 = [1, 3, 5, 6, 7, 10]
// nums2 = [2, 4, 5, 9, 11]

// 上图这两个给定数组A和B，一个长度是6，一个长度是5，归并之后的大数组仍然要保持升序，结果如下：
// [1, 2, 3, 4, 5, 5, 6, 7, 9, 10, 11]
// 大数组的长度是奇数（11），中位数显然是位于正中的第6个元素，也就是元素5。
// 中位数 5

// 2.偶数
// nums1 = [3, 5, 7, 8, 9]
// nums2 = [1, 2, 6, 7, 15]
// [1, 2, 3, 5, 6, 7, 7, 8, 9, 15]
// 中位数 (6 + 7)/2 = 6.5

// 解题思路：合并数组 快速排序，根据长度分奇偶，然后取值计算

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	tmp := sort.QuickSort(append(nums1,nums2...))
	len := len(tmp)

	if 0 == (len % 2) { // 偶数中位数
		return EvenNumbers(tmp,len/2)
	}
	return OddNumber(tmp,len/2)
	
}


// 奇数
func OddNumber(nums []int, k int) float64 {
	return float64(nums [k])
}

// 偶数
func EvenNumbers(nums []int,k int) float64{
	return (float64(nums [k-1])+float64(nums [k]))/2
}

