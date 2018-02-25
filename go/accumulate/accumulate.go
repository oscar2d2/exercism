package accumulate

func Accumulate(input []string, f func(string) string) (output []string) {
	for _, v := range input {
		output = append(output, f(v))
	}
	return
}
