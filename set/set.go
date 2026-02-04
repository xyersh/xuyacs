package set

// Set is a generic collection of unique elements
type Set[T comparable] struct {
	values map[T]struct{}
}

// NewSet creates a new empty set
func New[T comparable]() *Set[T] {

	return &Set[T]{
		values: make(map[T]struct{}),
	}
}

// Add adds an element to the set if it's not already present
func (s *Set[T]) Add(value T) {
	s.values[value] = struct{}{}
}

// Remove removes an element from the set if it exists
func (s *Set[T]) Remove(value T) {
	delete(s.values, value)
}

// Contains returns true if the set contains the given element
func (s *Set[T]) Contains(value T) bool {
	_, ok := s.values[value]
	return ok
}

// Size returns the number of elements in the set
func (s *Set[T]) Size() int {

	return len(s.values)
}

// Elements returns a slice containing all elements in the set
func (s *Set[T]) Elements() []T {
	result := make([]T, 0, len(s.values))
	for key, _ := range s.values {
		result = append(result, key)
	}
	return result
}
