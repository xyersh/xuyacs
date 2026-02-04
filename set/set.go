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

// Union returns a new set containing all elements from both sets
func Union[T comparable](s1, s2 *Set[T]) *Set[T] {
	result := New[T]()
	for key, _ := range s1.values {
		result.Add(key)
	}
	for key, _ := range s2.values {
		result.Add(key)
	}
	return result
}

// Intersection returns a new set containing only elements that exist in both sets
func Intersect[T comparable](s1, s2 *Set[T]) *Set[T] {
	result := New[T]()
	for key, _ := range s1.values {
		if s2.Contains(key) {
			result.Add(key)
		}
	}
	return result
}

// Difference returns a new set with elements in s1 that are not in s2
func Diff[T comparable](s1, s2 *Set[T]) *Set[T] {
	result := New[T]()
	for key, _ := range s1.values {
		if !s2.Contains(key) {
			result.Add(key)
		}
	}
	return result
}
