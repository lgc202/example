package double_pointer

//给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
//
// 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
//
// 返回容器可以储存的最大水量。
//
// 说明：你不能倾斜容器。
//
//
//
// 示例 1：
//
//
//
//
//输入：[1,8,6,2,5,4,8,3,7]
//输出：49
//解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
//
// 示例 2：
//
//
//输入：height = [1,1]
//输出：1
//
//
//
//
// 提示：
//
//
// n == height.length
// 2 <= n <= 10⁵
// 0 <= height[i] <= 10⁴
//
//

// 思路：使用对撞双指针
func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	maxArea := (right - left) * getMinHeight(height[left], height[right])
	for left < right {
		if height[left] < height[right] {
			// 面积 = x * y
			// 如果左边的高度更小, 就往右边找更高的才有可能得到更大的面积
			// 如果左边的高度已经比右边高了，那往右移动也没用, 面积反而变得更小, 因为y轴不变, x轴变小
			left++
		} else {
			right--
		}

		area := (right - left) * getMinHeight(height[left], height[right])
		maxArea = getMaxArea(maxArea, area)
	}

	return maxArea
}

func getMinHeight(h1, h2 int) int {
	if h1 < h2 {
		return h1
	}

	return h2
}

func getMaxArea(a1, a2 int) int {
	if a1 > a2 {
		return a1
	}

	return a2
}
