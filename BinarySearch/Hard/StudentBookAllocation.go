package hard

func findPages(nums []int, m int) int {
	start := 0
	end := 0
	for _, num := range nums {
		if num > start {
			start = num
		}
		end += num
	}
	res := -1
	for start <= end {
		mid := (start + end) / 2
		no := checkNoOfStudents(nums, mid)
		if no <= m {
			res = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return res
}
func checkNoOfStudents(nums []int, maxPages int) int {
	noOfStudents := 1
	pages := 0
	for _, num := range nums {
		if pages+num > maxPages {
			noOfStudents++
			pages = 0
		} else {
			pages += num
		}
	}
	return noOfStudents
}
