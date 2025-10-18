package basic

func getFloorAndCeil(nums []int, x int) []int {
	return []int{
		binaryGetFloor(nums, 0, len(nums)-1, x),
		binaryGetCeil(nums, 0, len(nums)-1, x),
	}
}

func binaryGetFloor(nums []int, start, end, target int) int {
	floor := -1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return nums[mid]
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			floor = nums[mid]
			start = mid + 1
		}
	}
	return floor
}

func binaryGetCeil(nums []int, start, end, target int) int {
	ceil := -1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return nums[mid]
		} else if nums[mid] > target {
			ceil = nums[mid]
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return ceil
}
