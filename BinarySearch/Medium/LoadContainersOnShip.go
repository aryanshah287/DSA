package medium

func shipWithinDays(weights []int, days int) int {
	start := 1
	end := 0
	for _, weight := range weights {
		end += weight
		if weight > start {
			start = weight
		}
	}
	ans := 0
	for start <= end {
		mid := (start + end) / 2
		d := checkCapacity(weights, mid)
		if d <= days {
			ans = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return ans
}
func checkCapacity(weights []int, cap int) int {
	days := 1
	sum := 0
	for _, weight := range weights {
		if (sum + weight) > cap {
			days++
			sum = weight
		} else {
			sum += weight
		}
	}
	return days
}
