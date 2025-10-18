package medium

import "math"

func minEatingSpeed(piles []int, h int) int {
	max := math.MinInt
	for _, n := range piles {
		if max < n {
			max = n
		}
	}
	start := 1
	end := max
	ans := max
	for start <= end {
		mid := (start + end) / 2
		totalHrs := callI(piles, mid)
		if totalHrs <= h {
			ans = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return ans
}
func callI(piles []int, k int) int {
	ans := 0
	for _, n := range piles {
		ans += (n + k - 1) / k
	}
	return ans
}
