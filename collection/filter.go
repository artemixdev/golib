package collection

func Filter[T any](input []T, callback func(index int, element T) bool) []T {
	output := make([]T, 0, len(input))
	for index, element := range input {
		if callback(index, element) {
			output = append(output, element)
		}
	}
	return output
}
