package valid_parentheses

/*
* test
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
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := []rune{}

	for _, char := range s {
		if (char == '(' || char == '{' || char == '[') {
			stack = append(stack, char)
		} else if (len(stack) > 0 && stack[len(stack) - 1] == chars[char]) {
			stack = stack[:len(stack) - 1]
		} else {
			return false, nil
		}
	}

	return len(stack) == 0, nil
}