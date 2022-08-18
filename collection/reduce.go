package collection

func Reduce[T, U any](input []T, initial U, callback func(accumulator U, element T) U) U {
	accumulator := initial
	for _, element := range input {
		accumulator = callback(accumulator, element)
	}
	return accumulator
}
