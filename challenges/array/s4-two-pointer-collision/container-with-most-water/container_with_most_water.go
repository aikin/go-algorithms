package container_with_most_water

import "errors"

/*
Given height = [1, 8, 6, 2, 5, 4, 8, 3, 7]

Because min(8,7)∗7 = 49

return 49

复杂度分析：
 - 时间复杂度：O(n^2)
 - 空间复杂度：O(1)，只需要常数级别的存储
*/

func MaxAreaByBruteForceWay(height []int) (int, error) {
	maxArea := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			maxArea = max(maxArea, calcArea(height, i, j))
		}
	}
	return maxArea, nil;
}

/*
* Coding
* 校验：
		-	非负数，
		-  2=< len(nums) <=  3 * 104
* 中间变量：双指针，分别指向 最左边 和 最右边
* 遍历：one for loop
* 操作：
		-	根据面积公式，计算面积。
		-	看左、右两边哪边是短，让短的一边向前移一步，看是否能获得更大的面积
* 复杂度分析：
* 	- 时间复杂度：O(n)
* 	- 空间复杂度：O(1)
*/
func MaxAreaByTwoPointerWay(height []int) (int, error) {
	if len(height) < 2 {
		return 0, errors.New("should >=2")
	}

	left, right := 0, len(height)-1

	maxArea := 0

	for left < right {
		maxArea = max(maxArea, calcArea(height, left, right))
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}

	return maxArea, nil
}

func calcArea(height []int, left int, right int) int {
	return min(height[left], height[right]) * (right - left)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
