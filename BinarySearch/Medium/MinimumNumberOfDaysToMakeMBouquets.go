package hard

func minDays(bloomDay []int, m int, k int) int {
	ans := -1
	start := 1
	end := func(b []int) int {
		max := 0
		for _, day := range bloomDay {
			if day > max {
				max = day
			}
		}
		return max
	}(bloomDay)
	for start <= end {
		mid := (start + end) / 2
		bouquets := calculateBouquets(bloomDay, mid, k)
		if bouquets >= m {
			ans = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return ans
}
func calculateBouquets(bloomDay []int, days int, k int) int {
	noOfBouquets := 0
	count := 0
	for _, day := range bloomDay {
		if days >= day {
			count++
			if count == k {
				noOfBouquets++
				count = 0
			}
		} else {
			count = 0
		}

	}
	return noOfBouquets
}
