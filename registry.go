package reg

import (
	"sync"
)

type Registry[T any] struct {
	values map[string]T
	mutex sync.Mutex
}

func (receiver *Registry[T]) Get(name string) (value T, found bool) {
	if nil == receiver {
		panic(errNilReceiver)
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	if nil == receiver.values {
		var empty T
		return empty, false
	}

	value, found = receiver.values[name]
	return value, found
}

func (receiver *Registry[T]) Len() int {
	if nil == receiver {
		panic(errNilReceiver)
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	return len(receiver.values)
}

func (receiver *Registry[T]) Set(name string, value T) (previous T, found bool) {
	if nil == receiver {
		panic(errNilReceiver)
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	if nil == receiver.values {
		receiver.values = map[string]T{}
	}

	previous, found = receiver.values[name]

	receiver.values[name] = value
	return previous, found
}

func (receiver *Registry[T]) Unset(name string) {
	if nil == receiver {
		panic(errNilReceiver)
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	if nil == receiver.values {
		return
	}

	delete(receiver.values, name)
}
