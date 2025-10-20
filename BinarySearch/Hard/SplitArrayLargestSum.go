package hard

func splitArray(nums []int, k int) int {
	start := 0
	end := 0
	for _, num := range nums {
		if num > start {
			start = num
		}
		end += num
	}
	ans := -1
	for start <= end {
		mid := (start + end) / 2
		splits := checkSplits(nums, mid)
		if splits <= k {
			ans = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return ans
}
func checkSplits(nums []int, size int) int {
	sum := 0
	count := 1
	for _, num := range nums {
		if sum+num <= size {
			sum += num
		} else {
			count++
			sum = num
		}
	}
	return count
}
