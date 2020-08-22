package group_anagrams

import (
	"sort"
	"strings"
)

/*
Given anagrams = ["eat", "tea", "tan", "ate", "nat", "bat"]

Because group anagrams [e, a, t] = [t, e, a]
Then [
	["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]
notes:
- 所有输入均为小写字母。
- 不考虑答案输出的顺序。

复杂度分析：
 - 时间复杂度：O(1)
 - 空间复杂度：O(1)
*/
func GroupAnagrams(strs []string) [][]string {

	ans := make(map[string][]string)

	for i := 0; i < len(strs); i++ {
		key := keyOf(strs[i])
		value, ok := ans[key]
		if (ok) {
			ans[key] = append(value, strs[i])
			continue
		}
		ans[key] = []string {strs[i]}
	}
	return values(ans)
}

func keyOf(s string) string {
	sChars := strings.Split(s, "")
	sort.Strings(sChars)
	return strings.Join(sChars, "")
}

func values(m map[string][]string) [][]string {
	v := make([][]string, 0, len(m))
	for  _, value := range m {
	 v = append(v, value)
	}
	return v
}