package collection

func Reduce[T, U any](initial U, input []T, callback func(accumulator U, element T) U) U {
	accumulator := initial
	for _, element := range input {
		accumulator = callback(accumulator, element)
	}
	return accumulator
}
