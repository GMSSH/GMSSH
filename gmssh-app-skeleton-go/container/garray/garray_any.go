package garray

import (
	"reflect"
	"sync"
)

// Array is a generic type alias for a slice of type T
type Array[T any] []T

// AnyArray is a thread-safe wrapper around a generic array/slice.
// It provides synchronized operations for concurrent access.
// Note: For single-threaded scenarios, using native Go slices is recommended for better performance.
type AnyArray[T any] struct {
	mu    sync.RWMutex // Mutex to protect concurrent access
	array []T          // Underlying slice storing the elements
}

// NewDefaultArray creates and returns a new empty AnyArray instance.
// The underlying slice is initialized with zero length and capacity.
func NewDefaultArray[T any]() *AnyArray[T] {
	return &AnyArray[T]{
		array: make([]T, 0),
	}
}

// NewArray creates and returns a new AnyArray initialized with the given data.
// It makes a copy of the input slice to ensure isolation.
func NewArray[T any](data []T) *AnyArray[T] {
	array := &AnyArray[T]{
		array: make([]T, 0),
	}
	array.array = append(array.array, data...)
	return array
}

// Append adds an element to the end of the array.
// This operation is thread-safe.
func (s *AnyArray[T]) Append(t T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.array = append(s.array, t)
}

// Index returns the element at the specified position.
// Panics if index is out of bounds (same behavior as native slices).
// This operation is thread-safe for reading.
func (s *AnyArray[T]) Index(i int) T {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.array[i]
}

// Last returns the last element of the array.
// Panics if array is empty.
// This operation is thread-safe for reading.
func (s *AnyArray[T]) Last() T {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.array[len(s.array)-1]
}

// First returns the first element of the array.
// Panics if array is empty.
// This operation is thread-safe for reading.
func (s *AnyArray[T]) First() T {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.array[0]
}

// Contains checks if the array includes the given element.
// Uses reflect.DeepEqual for comparison to handle complex types.
// This operation is thread-safe for reading.
func (s *AnyArray[T]) Contains(t T) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, v := range s.array {
		if reflect.DeepEqual(v, t) {
			return true
		}
	}
	return false
}

// Array returns a copy of the underlying slice.
// This operation is thread-safe for reading.
func (s *AnyArray[T]) Array() []T {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.array
}

// Reverse returns a new slice with elements in reverse order.
// The original array remains unchanged.
// This operation is thread-safe for reading.
func (s *AnyArray[T]) Reverse() []T {
	s.mu.RLock()
	defer s.mu.RUnlock()
	length := len(s.array)
	result := make([]T, length)
	for i, v := range s.array {
		result[length-1-i] = v // Place elements in reverse order
	}
	return result
}
