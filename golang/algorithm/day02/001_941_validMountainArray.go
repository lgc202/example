package day02

//给定一个整数数组 arr，如果它是有效的山脉数组就返回 true，否则返回 false。
//
// 让我们回顾一下，如果 arr 满足下述条件，那么它是一个山脉数组：
//
//
// arr.length >= 3
// 在 0 < i < arr.length - 1 条件下，存在 i 使得：
//
// arr[0] < arr[1] < ... arr[i-1] < arr[i]
// arr[i] > arr[i+1] > ... > arr[arr.length - 1]
//
//
//
//
//
//
//
//
//
// 示例 1：
//
//
//输入：arr = [2,1]
//输出：false
//
//
// 示例 2：
//
//
//输入：arr = [3,5,5]
//输出：false
//
//
// 示例 3：
//
//
//输入：arr = [0,3,2,1]
//输出：true
//
//
//
// 提示：
//
//
// 1 <= arr.length <= 10⁴
// 0 <= arr[i] <= 10⁴
//
//
// Related Topics 数组 👍 214 👎 0

// 本题跟山脉数组相关
func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	index := 0
	maxNum := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > maxNum {
			index = i
			maxNum = arr[i]
		}
	}

	// 如果最大值在第一个或最后一个元素, 直接返回false
	if index+1 == len(arr) || index == 0 {
		return false
	}

	// 检查上升
	for i := 1; i <= index; i++ {
		if arr[i-1] >= arr[i] {
			return false
		}
	}

	// 检查下降
	for i := index + 1; i < len(arr); i++ {
		if arr[i-1] <= arr[i] {
			return false
		}
	}

	return true
}

func validMountainArray2(arr []int) bool {
	i := 0
	n := len(arr)

	// 寻找最高点
	for i < n-1 && arr[i] < arr[i+1] {
		i++
	}

	// 最高点不能是第一个元素或者最后一个元素
	if i == 0 || i == n-1 {
		return false
	}

	// 从最高点往下降
	for i < n-1 && arr[i] > arr[i+1] {
		i++
	}

	// 到达最后一个元素
	return i == n-1
}
