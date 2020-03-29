package calculation

import (
	"sort"
)


// nums中找 a+b+c=0，要找出所有符合条件但 [不重复](也就是相同元素不可复用，但没有要求重复元素取哪一个) 的三元组

// a+b+c=0 <=> a+b=-c
// 相当于变形的两数之和问题

// 子问题
// 1.去除重复组合

// 思路：
// 排序 + 双指针
// 1. nums从小到大排序，i,j,k分别代表元素a,b,c的索引
// 2. 先固定i,i的范围是区间[0, len(nums)-2)，注意要对a去重，如果nums[i] == nums[i-1]将可能出现三元组重复
// 3. 固定i后，取j = i+1,k =len(nums)-1，j和k在j<k的情况下（j>i和j<k已经保证j和k在移动过程中不会数组越界）作为双指针进行遍历，判断a+b+c

func ThreeSum(nums []int) [][]int {
	
	if nums == nil || len(nums) < 3 {
		return nil
	}
	
	// 对nums升序重排
	sort.Ints(nums)
	len := len(nums)
	res := [][]int{}
	
	// 从第一个数开始往后匹配，留下最后两位。
	for i := 0; i < len-2; i++ {
	
		// 如果不是第一个需要 做去重处理
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		
		// 双指针方式
		l := i + 1    // j
		r := len - 1  // k
		
		// l >= r 当前的所有组合方式都已经匹配完成
		for l < r {
	
			s := nums[i] + nums[l] + nums[r]
		
			// 如果小于0 从左向右移动
			if s < 0 {
				l++
			}
			
			// 如果小于0 从右向左移动
			if s > 0 {
				r--
			}
			
			if s == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				
				// 下面两个循环为了出现相同的数，做去重处理
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				
				// 找到下一个组合
				l++
				r--
			}
			
		}
	}
	return res
}