package hard

func paint(A int, B int, C []int) int {
	start := 0
	end := 0
	for _, num := range C {
		if start < num {
			start = num
		}
		end += num
	}
	ans := -1
	for start <= end {
		mid := (start + end) / 2
		painters := checkNoOfPainters(C, mid)
		if painters <= A {
			ans = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return (ans * B) % 10000003
}
func checkNoOfPainters(units []int, maxTime int) int {
	size := 0
	painters := 1
	for _, unit := range units {
		if size+unit <= maxTime {
			size += unit
		} else {
			size = unit
			painters++
		}
	}
	return painters
}
