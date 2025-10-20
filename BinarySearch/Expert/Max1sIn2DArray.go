package expert

func rowWithMax1s(mat [][]int) int {
	rowIndex := -1
	maxCount := 0
	for i := 0; i < len(mat); i++ {
		countI := applyBinarySearch(mat[i])
		if countI > maxCount {
			maxCount = countI
			rowIndex = i
		}
	}
	return rowIndex
}

func applyBinarySearch(row []int) int {
	start := 0
	end := len(row) - 1
	lb := len(row)
	for start <= end {
		mid := (start + end) / 2
		if row[mid] == 1 {
			lb = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return len(row) - lb
}
