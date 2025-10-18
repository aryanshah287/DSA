package medium

func countOccurrences(nums []int, target int) int {
	first := findFirstOccurrenceCount(nums, target)
	if first == -1 {
		return 0 // target not found
	}
	last := findLastOccurrenceCount(nums, target)
	return last - first + 1
}

func findFirstOccurrenceCount(nums []int, target int) int {
	start, end := 0, len(nums)-1
	first := -1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			first = mid
			end = mid - 1
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return first
}

func findLastOccurrenceCount(nums []int, target int) int {
	start, end := 0, len(nums)-1
	last := -1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			last = mid
			start = mid + 1
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return last
}
