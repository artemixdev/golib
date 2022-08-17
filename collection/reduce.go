package collection

func Reduce[T, A any](input []T, initial A, callback func(accumulator A, element T) A) A {
	accumulator := initial
	for _, element := range input {
		accumulator = callback(accumulator, element)
	}
	return accumulator
}
