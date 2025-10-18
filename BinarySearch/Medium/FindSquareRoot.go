package medium

func floorSqrt(n int64) int64 {
	if n == 0 || n == 1 {
		return n
	}

	start := int64(1)
	end := n
	var ans int64

	for start <= end {
		mid := (start + end) / 2
		if mid <= n/mid {
			ans = mid
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return ans
}
