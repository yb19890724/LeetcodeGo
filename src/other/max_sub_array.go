package other

//给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

//输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
//输出：6
//解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。

// 思路：临时变量记录最大值，如果出现新的进行更新
func MaxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	maxSum, tmp, p := nums[0], 0, 0
	for p < len(nums) {

		tmp += nums[p]

		tmp, maxSum = maxSub(tmp, maxSum)

		p++
	}
	return maxSum
}

func maxSub(tmp int, maxSum int) (int, int) {

	if tmp > maxSum {
		maxSum = tmp
	}

	if tmp < 0 {
		tmp = 0
	}
	return tmp, maxSum
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// dp 算法
func DpMaxSubArray(nums []int) int {

	l:= len(nums)

	if 0 == l {
		return 0
	}
	if 1 == l {
		return nums[0]
	}

	dp, res := make([]int, l), nums[0]

	dp[0] = nums[0]

	for i := 1; i < l; i++ {
		if dp[i-1] > 0 {
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i]
		}
		res = max(res, dp[i])
	}
	return res
}
