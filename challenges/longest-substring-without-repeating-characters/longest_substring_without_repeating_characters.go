package longest_substring_without_repeating_characters

/*
GIVEN: s = "abcabcbb"
WHEN: The answer is "abc", with the length of 3.
THEN: 3

复杂度分析：
 - 时间复杂度：O(n2)
 - 空间复杂度：O(n)
*/
func LengthOfLongestSubstringByBruteForceWay(s string) int {
	if (len(s) == 0) {
		return 0
	}

	maxLen := 0
	len := len(s)

	getlong := func (l int) int {
		counter := make(map[byte]bool)
		for i:= l; i < len; i++{
				if counter[s[i]] {
						return i - l
				}
				counter[s[i]] = true
		}
		return len - l
	}

	for i := 0; i < len; i++ {
		if (maxLen > len - i) {
			break
		}
		longest := getlong(i)
		if longest > maxLen {
			maxLen = longest
		}
	}

	return maxLen
}