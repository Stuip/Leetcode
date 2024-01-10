package main

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}

	i, j, k := len(nums)-2, len(nums)-1, len(nums)-1

	// find: A[i]<A[j]
	for i >= 0 && nums[i] >= nums[j] {
		// 这个操作能确保[j, end]是降序的
		i--
		j--
	}

	if i >= 0 { // 不是最后一个排列
		// 在[j, end]中，寻找到第一个比nums[i]大的数值
		for nums[i] >= nums[k] {
			k--
		}
		// 交换 nums[i], nums[k]
		nums[i], nums[k] = nums[k], nums[i]
	}

	// 反转 A[j:end]
	for i, j := j, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func main() {
	nums := []int{3, 2, 1}
	nextPermutation(nums)
}
