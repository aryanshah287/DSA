package easy

func lowerBound(nums []int, target int) int {
	return lowerBoundRecursive(nums, 0, len(nums)-1, target, len(nums))
}

func lowerBoundRecursive(nums []int, low, high, target, ans int) int {
	if low > high {
		return ans
	}

	mid := (low + high) / 2

	if nums[mid] >= target {
		return lowerBoundRecursive(nums, low, mid-1, target, mid)
	}
	return lowerBoundRecursive(nums, mid+1, high, target, ans)
}
