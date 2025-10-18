package medium

func singleNonDuplicate(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if nums[0] != nums[1] {
		return nums[0]
	}
	if nums[len(nums)-2] != nums[len(nums)-1] {
		return nums[len(nums)-1]
	}
	start := 1
	end := len(nums) - 2
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] != nums[mid+1] && nums[mid] != nums[mid-1] {
			return nums[mid]
		}
		if (mid%2 == 0 && nums[mid] == nums[mid+1]) || (mid%2 == 1 && nums[mid] == nums[mid-1]) {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}
