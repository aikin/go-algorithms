package longest_repeating_character_replacement


/*
GIVEN: s = "ABAB", k = 2
WHEN: Replace the two 'A's with two 'B's or vice versa.
THEN: 4

复杂度分析：
 - 时间复杂度：O(n)
 - 空间复杂度：O(1)
*/

func CharacterReplacement(s string, k int) int {
	n := len(s)
	if (k < 0) {
		k = 0
	}
	if (k + 1 >= n) {
		return n
	}


	counter := make([]int, 26 + 'A')

	left, curMax, res := 0, 0, 0

	for right := 0; right < n; right++ {
		counter[s[right]]++

		if (counter[s[right]] > curMax) {
			curMax = counter[s[right]]
		}

		for right - left + 1 > curMax + k {
			counter[s[left]]--
			left++
		}

		if right - left + 1 > res {
			 res = right - left + 1
		}
	}
	return res
}