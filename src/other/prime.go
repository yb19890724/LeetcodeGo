package other

// 求素数
func Prime(value int) bool {

	// 小于等于3或者大于等于2
	if value <= 3 {
		return value >= 2
	}

	// 整除2和3都不是素数
	if value%2 == 0 || value%3 == 0 {
		return false
	}

	// 匹配整除5，或者7，以此累加
	// 例子特殊：11,17,23,29,35,41,47,53,59,65
	for i := 5; i*i <= value; i += 6 {

		if value%i == 0 || value%(i+2) == 0 {

			return false

		}

	}

	return true
}
