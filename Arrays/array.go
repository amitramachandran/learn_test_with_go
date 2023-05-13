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

	for _, sli := range slicess {
		result = append(result, SumOfSlice(sli[1:]))
	}
	return
}
