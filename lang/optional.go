package lang

type Optional[T any] struct {
	value   T
	present bool
}

func NewOptional[T any](value ...T) Optional[T] {
	if len(value) == 0 {
		return Optional[T]{}
	}
	return Optional[T]{value: value[0], present: true}
}

func (opt Optional[T]) Value() (value T, present bool) {
	return opt.value, opt.present
}
