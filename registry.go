package reg

import (
	"sync"
)

type Registry[T any] struct {
	values map[string]T
	mutex sync.Mutex
}

// For lets you iterate through all the items in the registry — it calls func 'fn' on each item in the registry.
//
// Note that you should NOT call .Get(), .Set(), .Len(), or .Unset() from the `fn`.
// It will cause For to lock.
func (receiver *Registry[T]) For(fn func(Registerer[T], string, T)) {
	if nil == receiver {
		panic(errNilReceiver)
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	var registerer Registerer[T] = RegistererFuncs[T]{
		GetFunc: receiver.get,
		LenFunc: receiver.len,
		SetFunc: receiver.set,
		UnsetFunc: receiver.unset,
	}

	for name, value := range receiver.values {
		fn(registerer, name, value)
	}
}

// Get return the item inthe registry registered under the name 'name'.
//
// This should NOT be called from within the function passed to [Registry.For].
// Doing so will cause a deadlock
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
//
// This should NOT be called from within the function passed to [Registry.For].
// Doing so will cause a deadlock
func (receiver *Registry[T]) Len() int {
	if nil == receiver {
		panic(errNilReceiver)
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	return receiver.len()
}


func (receiver *Registry[T]) len() int {
	return len(receiver.values)
}

// Set registers an item in the registry under the name 'name', but
// it also returns the previous item under the name 'name' if it existed.
//
// This should NOT be called from within the function passed to [Registry.For].
// Doing so will cause a deadlock
func (receiver *Registry[T]) Set(name string, value T) (previous T, found bool) {
	if nil == receiver {
		panic(errNilReceiver)
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	return receiver.set(name, value)
}

func (receiver *Registry[T]) set(name string, value T) (previous T, found bool) {
	if nil == receiver.values {
		receiver.values = map[string]T{}
	}

	previous, found = receiver.get(name)

	receiver.values[name] = value
	return previous, found
}

// Unset removes an item in the registry under the name 'name', if it is there, and
// it also returns the previous item under the name 'name' if it existed.
//
// This should NOT be called from within the function passed to [Registry.For].
// Doing so will cause a deadlock
func (receiver *Registry[T]) Unset(name string) (previous T, found bool) {
	if nil == receiver {
		panic(errNilReceiver)
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	return receiver.unset(name)
}

func (receiver *Registry[T]) unset(name string) (previous T, found bool) {
	if nil == receiver.values {
		return
	}

	previous, found = receiver.get(name)

	delete(receiver.values, name)
	return previous, found
}


// UnsetWhen removes an item in the registry under the name 'name', if it is there and the `whenFunc` returns true, and
// it also returns the previous item under the name 'name' if it existed.
//
// This should NOT be called from within the function passed to [Registry.For].
// Doing so will cause a deadlock
func (receiver *Registry[T]) UnsetWhen(name string, whenFunc func(T)bool) (previous T, found bool, when bool) {
	if nil == receiver {
		panic(errNilReceiver)
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	{
		previous, found = receiver.get(name)
		if found && !whenFunc(previous) {
			return
		}
	}

	{
		previous, found := receiver.unset(name)
		return previous, found, true
	}
}
