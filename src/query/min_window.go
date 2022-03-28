package query

import "fmt"

// @todo 要求
//给定一个源字符串 s，再给一个字符串 T，
//要求在源字符串中找到一个窗口，这个窗口包含由字符串各种排列组合组成的，窗口中可以包含 T 中没有的字符，如果存在多个，
//在结果中输出最小的窗口，如果找不到这样的窗口，输出空字符串。

//最小覆盖子串是这样的子串，不仅字符串t中的每个字符都在该子串中出现，同时该子串中那些和字符串t中字符一样的字符其出现次数要大于等于其在字符串t中出现的次数。该题目用滑动窗口思想来解决的整体思路是：

//首先，统计字符串t中每个字符的出现次数。

//然后，扩大窗口右侧边界，直到窗口内的字符串包含了字符串t中的每个字符，同时字符串t中的每个字符在窗口内的字符串中的出现次数要大于等于字符串t中每个字符的出现次数。

//接着，尝试缩小窗口的左侧边界，看窗口内字符串能不能覆盖字符串t。


func MinWindow(s string, t string) string {
	if s == "" || t == "" {
		return ""
	}

	var (
		// tFreq 统计字符串t中每个字符的出现次数
		tFreq [256]int
		// wFreq 用于记录窗口中每个字符的出现次数
		windowFreq [256]int
		// windowSameCount 当前窗口中(s字符串)和字符串t中的字符相同的字符个数  更新
		windowSameCount = 0

		// 开始位置
		start = 0
        // 结束位置
		end = -1

		// 最小窗口的最左边
		startIndex = -1
		// 最小窗口的最右边
		minWindowLength = len(s) + 1
	)

	// 记录t每个字符串出现的次数
	for i := 0; i < len(t); i++ {
		tFreq[t[i]]++
	}

	// 移动窗口
	for start < len(s) {

		if end+1 < len(s) && windowSameCount < len(t) {
			windowFreq[s[end+1]]++
			// 窗口中当前考察的字符的出现次数小于等于字符串t中字符出现次数
			// 则窗口中当前考察的字符是在字符串t中出现的
			if windowFreq[s[end+1]] <= tFreq[s[end+1]] {
				windowSameCount++
			}
			end++
			continue
		}

		// 字符串t中的字符在窗口中全部出现
		if windowSameCount == len(t) && end-start+1 < minWindowLength {
			minWindowLength = end - start + 1
			startIndex = start
			fmt.Println(s[start : end+1])
		}
		// 窗口左移
		windowFreq[s[start]]--
		if windowFreq[s[start]] < tFreq[s[start]] {
			windowSameCount--
		}
		start++
	}

	if startIndex != -1 {
		return s[startIndex : startIndex+minWindowLength]
	}

	return ""
}
