package calculation

import (
	"fmt"
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


	if nil == nums || len(nums) < 3  {
		return nil
	}

	sort.Ints(nums)

	m := make(map[string]string)
	count :=len(nums)
	res := [][]int{}

	for i:=0;i < count -2; i++ {

		// 双指针方式
		l := i + 1    // l
		r := count - 1  // r

		for l < r {

			s := nums[i] + nums[l] + nums[r]

			if s > 0 {
				r--
			}

			if s < 0 {
				l++
			}

			if 0 == s {

				tmpKey:= fmt.Sprintf("%d%d%d",i,l,r)
				if _,ok :=m[tmpKey]; !ok {
					res = append(res, []int{nums[i],nums[l],nums[r]})
					m[tmpKey] =tmpKey
				}

				l++
				r--
			}

		}
	}

	return res
}