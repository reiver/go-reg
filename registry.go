package reg

import (
	"sync"
)

type Registry[T any] struct {
	values map[string]T
	mutex sync.Mutex
}

// For lets you iterate through all the items in the registry — it calls func 'fn' on each item in the registry.
func (receiver *Registry[T]) For(fn func(string, T)) {
	if nil == receiver {
		panic(errNilReceiver)
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	for name, value := range receiver.values {
		fn(name, value)
	}
}

// Get return the item inthe registry registered under the name 'name'.
func (receiver *Registry[T]) Get(name string) (value T, found bool) {
	if nil == receiver {
		panic(errNilReceiver)
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	return receiver.get(name)
}

func (receiver *Registry[T]) get(name string) (value T, found bool) {
	var empty T

	if nil == receiver {
		return empty, false
	}
	if nil == receiver.values {
		return empty, false
	}

	value, found = receiver.values[name]
	return value, found
}

// Len returns the number of items in the registry.
func (receiver *Registry[T]) Len() int {
	if nil == receiver {
		panic(errNilReceiver)
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	return len(receiver.values)
}

// Set registers an item in the registry under the name 'name', but
// it also returns the previous item under the name 'name' if it existed.
func (receiver *Registry[T]) Set(name string, value T) (previous T, found bool) {
	if nil == receiver {
		panic(errNilReceiver)
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	if nil == receiver.values {
		receiver.values = map[string]T{}
	}

	previous, found = receiver.get(name)

	receiver.values[name] = value
	return previous, found
}

// Unset removed an item in the registry under the name 'name', if it is there, and
// it also returns the previous item under the name 'name' if it existed.
func (receiver *Registry[T]) Unset(name string) (previous T, found bool) {
	if nil == receiver {
		panic(errNilReceiver)
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	if nil == receiver.values {
		return
	}

	previous, found = receiver.get(name)

	delete(receiver.values, name)
	return previous, found
}
