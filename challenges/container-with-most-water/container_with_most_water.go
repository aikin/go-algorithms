package container_with_most_water

/*
Given height = [1, 8, 6, 2, 5, 4, 8, 3, 7]

Becaue min(8,7)∗7 = 49

return 49

复杂度分析：
 - 时间复杂度：O(n^2)
 - 空间复杂度：O(1)，只需要常数级别的存储
*/

func MaxArea(height []int) (int, error) {
	maxArea := 0;
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			maxArea = max(maxArea, min(height[i], height[j]) * (j - i));
		}
	}
	return maxArea, nil;
}

func max(x, y int) int {
	if (x > y) {
		return x
	}
	return y
}

func min(x, y int) int {
	if (x > y) {
		return y
	}
	return x
}