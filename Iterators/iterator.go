package iterators

func Repeat(input string, times int) string {
	var result string
	for i := 0; i < times; i++ {
		result += input
	}
	return result
}
