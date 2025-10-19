package hard

import "sort"

func aggressiveCows(nums []int, k int) int {
	start := 1
	sort.Ints(nums)
	end := nums[len(nums)-1] - nums[0]
	res := -1
	for start <= end {
		mid := (start + end) / 2
		ans := checkIfFits(nums, k, mid)
		if ans == true {
			res = mid
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return res
}
func checkIfFits(nums []int, k int, dist int) bool {
	numOfCows := 1
	lastStall := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]-lastStall >= dist {
			lastStall = nums[i]
			numOfCows++
		}
	}
	return numOfCows >= k
}
