package arrays

func SumOfArray(a [5]int) int {
	var sum int
	for _, number := range a {
		sum += number
	}
	return sum
}

func SumOfSlice(a []int) int {
	var sum int
	for _, number := range a {
		sum += number
	}
	return sum
}

// variadic function with multiple args passable
func SumAll(slicess ...[]int) []int {
	var result []int

	for _, sli := range slicess {
		result = append(result, SumOfSlice(sli))
	}
	return result

}

func SumAllTails(slicess ...[]int) (result []int) {
	var sum_tail int
	for _, sli := range slicess {
		if len(sli) == 0 {
			sum_tail = 0
		} else {
			sum_tail = SumOfSlice(sli[1:])
		}
		result = append(result, sum_tail)
	}
	return
}
