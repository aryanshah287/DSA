package hard

func NthRoot(N int, M int) int {
	start := int64(0)
	end := int64(M)
	var ans int64 = -1

	for start <= end {
		mid := (start + end) / 2
		pow := int64(1)
		for i := 0; i < N; i++ {
			pow *= mid
			if pow > int64(M) {
				break
			}
		}

		if pow <= int64(M) {
			ans = mid
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return int(ans)
}
