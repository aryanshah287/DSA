package medium

func findKthPositive(arr []int, k int) int {
	start := 0
	end := len(arr) - 1
	for start <= end {
		mid := (start + end) / 2
		missing := arr[mid] - (mid + 1)
		if missing < k {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return end + k + 1
}
