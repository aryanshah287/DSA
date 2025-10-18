package easy

func searchRange(nums []int, target int) []int {
	var ans []int
	ans = append(ans, findFirstOccurrence(nums, target))
	ans = append(ans, findLastOccurrence(nums, target))
	return ans
}
func findLastOccurrence(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	last := -1
	if len(nums) == 0 {
		return last
	}
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
func findFirstOccurrence(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	last := -1
	if len(nums) == 0 {
		return last
	}
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			last = mid
			end = mid - 1
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return last
}
