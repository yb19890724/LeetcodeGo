package calculation

// 思路：判断数组中所有值得组合方式，是否满足条件

func TwoSum(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		
		for j := i + 1; j < len(nums); j++ {
			
			if (nums[i] + nums[j]) == target {
				
				return []int{i, j}
				
			}
		}
	}

	return []int{}
}