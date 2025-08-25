package reg

type RegistererFuncs[T any] struct {
	GetFunc   func(name string) (value T, found bool)
	LenFunc   func() int
	SetFunc   func (name string, value T) (previous T, found bool)
	UnsetFunc func(name string) (previous T, found bool)
}

var _ Registerer[uint64] = RegistererFuncs[uint64]{}

func (receiver RegistererFuncs[T]) Get(name string) (value T, found bool) {
	return receiver.GetFunc(name)
}

func (receiver RegistererFuncs[T]) Len() int {
	return receiver.LenFunc()
}

func (receiver RegistererFuncs[T]) Set(name string, value T) (previous T, found bool) {
	return receiver.SetFunc(name, value)
}

func (receiver RegistererFuncs[T]) Unset(name string) (previous T, found bool) {
	return receiver.UnsetFunc(name)
}
