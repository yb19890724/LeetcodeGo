package query

// 实现一个strstr 函数
// 给定一个haystack字符串和一个needle字符串，在haystack字符串中找出needle字符串出现的第一个位置(从0开始)。不存在返回-1


//实现一个查找 substring 的函数。如果在母串中找到了子串，返回子串在母串中出现的下标，如果没有找到，返回 -1，如果子串是空串，则返回 0 。
func StrStr(haystack string, needle string) int {

	for i := 0; ; i++ {
		for j := 0; ; j++ {

			// 空串 或者下标和查找的字符串长度相等那么返回下标
			if j == len(needle) {
				return i
			}

			// 没找到
			if i+j == len(haystack) {
				return -1
			}

			// 从0开始找，找到不相等的位置就跳出这个循环
			if needle[j] != haystack[i+j] {
				break
			}
		}

	}
}
