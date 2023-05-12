package iterators

func Repeat(input string) string {
	result := ""
	for i := 0; i < 5; i++ {
		result = result + input
	}
	return result
}
