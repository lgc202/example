package array

func canThreePartsEqualSum(arr []int) bool {
	sum := 0
	for i := range arr {
		sum += arr[i]
	}

	// 如果可以分成3部分,肯定可以被3整除
	if sum%3 != 0 {
		return false
	}

	// 每一部分的和
	key := sum / 3

	// 利用对撞双指针
	left := 0
	right := len(arr) - 1
	leftSum := arr[left]
	rightSum := arr[right]

	for left+1 < right { // 要保证中间部分至少有一个元素
		if leftSum == key && rightSum == key {
			return true
		}

		if leftSum != key {
			left++
			leftSum += arr[left]
		}

		if rightSum != key {
			right--
			rightSum += arr[right]
		}
	}

	return false
}
