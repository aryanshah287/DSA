package expert

func searchMatrix(matrix [][]int, target int) bool {
	start := 0
	end := len(matrix)*len(matrix[0]) - 1
	for start <= end {
		mid := (start + end) / 2
		row := mid / len(matrix[0])
		col := mid % len(matrix[0])
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return false
}
