package collection

func Reduce[T, U any](initial U, input []T, callback func(acc U, index int, element T) U) U {
	acc := initial
	for index, element := range input {
		acc = callback(acc, index, element)
	}
	return acc
}
