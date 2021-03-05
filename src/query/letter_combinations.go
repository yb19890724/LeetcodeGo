package query

var result []string
var dict = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func LetterCombinationsBT(digits string) []string {
	result = []string{}
	if digits == "" {
		return result
	}
	letter("", digits)
	return result
}

func letter(res string, digits string) {
	if digits == "" {
		result = append(result, res)
		return
	}

	k := digits[0:1]
	digits = digits[1:]

	for i := 0; i < len(dict[k]); i++ {
		res += dict[k][i]
		letter(res, digits)
		res = res[0 : len(res)-1]
	}
}
