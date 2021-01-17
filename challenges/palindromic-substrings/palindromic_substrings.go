package palindromic_substrings


func CountSubstringsByBruteForceWay(s string) int {
	sum := 0

	for i := 0; i < len(s); i++ {
		for j := i; j <= len(s); j++ {
			if isSubstrings(s[i:j]) {
				sum++
			}
		}
	}

	return sum
}

func isSubstrings(s string) bool {
	if s == "" {
		return false
	}

	l, r := 0, len(s) - 1

	for i := 0; i < len(s) / 2; i++ {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

