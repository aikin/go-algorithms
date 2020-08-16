package merge_intervals

import (
	"sort"
)

/*
Given nums = intervals = [[1, 3], [2, 6], [8, 10], [15, 18]]

Because 区间 [1, 3] 和 [2, 6] 重叠, 将它们合并为 [1, 6].
Then result: [[1, 6], [8, 10], [15, 18]]


复杂度分析：
 - 时间复杂度：O(nlogn)，更多的时间消耗，在排序算法
 - 空间复杂度：O(logn)

 贪心算法：可以适用贪心的问题就是每一步局部最优，最后导致结果全局最优。
 merged 就是每一步局部最优解的 存储容器。
*/
func Merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{}

	for _, curInterval := range intervals {
		if (len(merged) == 0) {
			merged = append(merged, curInterval)
			continue
		}
		last := merged[len(merged) - 1]
		if (last[1] < curInterval[0]) {
			merged = append(merged, curInterval)
			continue
		}
		last[1] = max(last[1], curInterval[1])
	}
	return merged
}

func max(x, y int) int {
	if (x > y) {
		return x
	}
	return y
}