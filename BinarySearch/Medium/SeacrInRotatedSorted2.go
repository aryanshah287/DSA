package medium

func search2(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	start := 0
	end := len(nums) - 1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return true
		}
		if nums[start] == nums[mid] && nums[end] == nums[mid] {
			start++
			end--
			continue
		}
		if nums[start] <= nums[mid] {
			if nums[start] <= target && target <= nums[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			if nums[mid] <= target && target <= nums[end] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return false
}
