package collection

func Filter[T any](input []T, callback func(index int, element T) bool) []T {
	output := make([]T, len(input))
	FilterTo[T](input, &output, callback)
	return output
}

func FilterTo[T any](input []T, output *[]T, callback func(index int, element T) bool) {
	outIdx := 0
	for inIdx, element := range input {
		if callback(inIdx, element) {
			(*output)[outIdx] = element
			outIdx++
		}
	}
	*output = (*output)[:outIdx]
}
