package merge_intervals

import (
	"sort"
)

/*
Given nums = intervals = [[1, 3], [2, 6], [8, 10], [15, 18]]

Because 区间 [1, 3] 和 [2, 6] 重叠, 将它们合并为 [1, 6].
Then result: [[1, 6], [8, 10], [15, 18]]


复杂度分析：
 - 时间复杂度：
 - 空间复杂度：
*/
func Merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{}

	for i := 0; i < len(intervals); i++ {
		if (len(merged) == 0) {
			merged = append(merged, intervals[i])
			continue
		}
		if (merged[len(merged) - 1][1] < intervals[i][0]) {
			merged = append(merged, intervals[i])
			continue
		}
		merged[len(merged) - 1][1] = max(merged[len(merged) - 1][1], intervals[i][1])
	}

	return merged
}

func max(x, y int) int {
	if (x > y) {
		return x
	}
	return y
}