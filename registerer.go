package reg

type Registerer[T any] interface {
	Get(name string) (value T, found bool)
	Len() int
	Set(name string, value T) (previous T, found bool)
	Unset(name string) (previous T, found bool)
}
