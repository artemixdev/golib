package collection

func Map[T, M any](input []T, callback func(index int, element T) M) []M {
	output := make([]M, 0, len(input))
	for index, element := range input {
		output = append(output, callback(index, element))
	}
	return output
}
