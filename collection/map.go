package collection

func Map[T, U any](input []T, callback func(index int, element T) U) []U {
	output := make([]U, len(input))
	MapTo[T, U](input, &output, callback)
	return output
}

func MapTo[T, U any](input []T, output *[]U, callback func(index int, element T) U) {
	for index, element := range input {
		(*output)[index] = callback(index, element)
	}
	*output = (*output)[:len(input)]
}
