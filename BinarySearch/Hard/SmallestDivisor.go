package hard

func smallestDivisor(nums []int, threshold int) int {
	start := 1
	ans := 0
	end := nums[0]
	for _, num := range nums {
		if num > end {
			end = num
		}
	}
	for start <= end {
		mid := start + (end-start)/2
		sum := calculateSum(nums, mid)
		if sum <= threshold {
			ans = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return ans
}
func calculateSum(nums []int, divisor int) int {
	sum := 0
	for _, dividend := range nums {
		sum += (dividend + divisor - 1) / divisor
	}
	return sum
}
