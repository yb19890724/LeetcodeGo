package main

// 滑动窗口模式
func lengthOfLongestSubstring(s string) int {
	m := map[rune]int{}

	start, max := -1, 0

	for index, v := range s {

		if last, ok := m[v]; ok && last > start {
			start = last
		}

		m[v] = index

		if index-start > max {
			max = index - start
		}
	}

	return max
}
