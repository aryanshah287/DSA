package easy

func searchInsert(nums []int, target int) int {
	return recursiveBinarySearchInsert(nums, 0, len(nums)-1, target, len(nums))
}
func recursiveBinarySearchInsert(nums []int, low int, high int, target int, ans int) int {
	if low > high {
		return ans
	}
	mid := (low + high) / 2
	if nums[mid] >= target {
		return recursiveBinarySearchInsert(nums, low, mid-1, target, mid)
	}
	return recursiveBinarySearchInsert(nums, mid+1, high, target, ans)
}
