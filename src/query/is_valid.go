package query

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

//有效字符串需满足：
//1.左括号必须用相同类型的右括号闭合。
//2.左括号必须以正确的顺序闭合。

// @todo 解题思路
//   遇到左括号就进栈push，遇到右括号并且栈顶为与之对应的左括号，就把栈顶元素出栈。最后看栈里面还有没有其他元素，如果为空，即匹配。
//   需要注意，空字符串是满足括号匹配的，即输出 true。
func IsValid(s string) bool {

	// 空字符串直接返回 true
	if len(s) == 0 {
		return true
	}

	stack := make([]rune, 0)

	for _, v := range s {

		// 匹配左括号
		if (v == '[') || (v == '(') || (v == '{') {
			stack = append(stack, v)
			continue
		}

		if (v == ']') && len(stack) > 0 && stack[len(stack)-1] == '[' {
			stack = stack[:len(stack)-1]
			continue
		}

		if (v == ')') && len(stack) > 0 && stack[len(stack)-1] == '(' {
			stack = stack[:len(stack)-1]
			continue
		}

		if (v == '}') && len(stack) > 0 && stack[len(stack)-1] == '{' {
			stack = stack[:len(stack)-1]
			continue
		}

		return false
	}

	return true
}
