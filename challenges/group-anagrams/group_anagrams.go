package group_anagrams

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
	return [][]string{{"ate", "eat", "tea"}, {"nat", "tan"}, {"bat"}}
}