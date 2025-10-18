package easy

func search(nums []int, target int) int {
	return recursiveBinary(nums, 0, len(nums)-1, target)
}

func recursiveBinary(nums []int, low int, high int, target int) int {
	if low > high {
		return -1
	}
	mid := (low + high) / 2
	if target == nums[mid] {
		return mid
	} else if target > nums[mid] {
		return recursiveBinary(nums, mid+1, high, target)
	}
	return recursiveBinary(nums, low, mid-1, target)
}
