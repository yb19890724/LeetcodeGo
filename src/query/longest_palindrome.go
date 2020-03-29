package query

// 给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
//
// 示例 1：
//
// 输入: "babad"
// 输出: "bab"
// 注意: "aba" 也是一个有效答案。
// 示例 2：
//
// 输入: "cbbd"
// 输出: "bb"

// 最长回文子串
// 找到两个字符串非连续最大公共字符串。
func LongestPalindrome(s string) string {
	
	str := []byte(s)
	l := len(str)
	if l == 0 {
		return ""
	}
	
	// 根据字符串的长度，建立一个二维切片，记录匹配到的结果
	dp:= dp(l)
	
	// 记录每一次找到最长匹配的长度和串
	ret := []byte{}
	max := 0
	
	// 匹配流程:
	// i是从右往左匹配
	// j是从左往右匹配
	
	// babad
	// 1. j=4,i=4 ret:d
	// 2. j=3,i=3 ret:d
	//    j=4,i=3 ret:d
	// 3. j=2,i=2 ret:d
	//    j=3,i=2 ret:d
	//    j=4,i=2 ret:d
	for i := l - 1; i >= 0; i-- {
		for j := i; j < l; j++ {
			
			// j-i <=1 左右匹配超过1个位置
			// dp[i+1][j-1] 兼容匹配左右俩那边是否满足
			if str[i] == str[j] && (j-i <=1 || dp[i+1][j-1]) {
				dp[i][j] = true
				// 记录 匹配到最长的字符串 更新
				if max < j-i+1 {
					max = j - i + 1
					ret = str[i : j+1]
				}
			}
			// fmt.Println(j-i,j,i)
			// fmt.Println(string(ret))
		}
		//fmt.Println("-------")
	}
	return string(ret)
}

// 动态规划解决的问题多数有重叠子问题这个特点，为减少重复计算，对每一个子问题只解一次，将其不同阶段的不同状态保存在一个二维切片中。
func dp(l int) [][]bool {
	dp := make([][]bool, l)
	for k := range dp {
		dp[k] = make([]bool, l)
	}
	return dp
}


