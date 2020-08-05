package valid_parentheses

/*
Given s = [{}]

Because is valid parentheses
return true

复杂度分析：
 - 时间复杂度：O(n)
 - 空间复杂度：O(n)，可能输入的 s 很长
*/
func IsValid(s string) (bool, error) {
	if (len(s) == 0) {
		return true, nil
	}

	invalidLen := len(s) % 2 != 0
	invalidStartChar := len(s) > 0 && (s[0] == '}' || s[0] == ')' || s[0] == ']')
	if (invalidLen || invalidStartChar) {
		return false, nil
	}

	chars := map[rune] rune {
		'(': ')',
		'{': '}',
		'[': ']',
	}

	stack := []rune{}

	for _, char := range s {
		if _, ok := chars[char]; ok {
			stack = append(stack, char)
			continue
		}
		topChar := stack[len(stack) - 1]
		matchedChar := chars[topChar]
		if (len(stack) > 0 && matchedChar == char) {
			stack = stack[:len(stack) - 1]
			continue
		}
		return false, nil
	}

	return len(stack) == 0, nil
}