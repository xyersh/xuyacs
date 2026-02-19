package queue

import "errors"

// ErrEmptyQueue is returned when an operation cannot be performed on an empty collection
var ErrEmptyQueue = errors.New("queue is empty")

type node[T any] struct {
	value T
	next  *node[T]
}

// Queue is a generic First-In-First-Out (FIFO) data structure
type Queue[T any] struct {
	head *node[T]
	tail *node[T]
	size int
}

// NewQueue creates a new empty queue
func New[T any]() *Queue[T] {

	return &Queue[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
}

// Enqueue adds an element to the end of the queue
func (q *Queue[T]) Enqueue(value T) {
	newNode := &node[T]{
		value: value,
		next:  nil,
	}

	if q.tail == nil {
		q.head = newNode
		q.tail = newNode
	} else {

		q.tail.next = newNode
		q.tail = newNode
	}
	q.size++
}

// Dequeue removes and returns the front element from the queue
// Returns an error if the queue is empty
func (q *Queue[T]) Dequeue() (T, error) {
	var result T
	if q.head == nil {
		return result, ErrEmptyQueue
	} else {
		result = q.head.value
		q.head = q.head.next
	}
	if q.head == nil {
		q.tail = nil
	}
	q.size--
	return result, nil

}

// Front returns the front element without removing it
// Returns an error if the queue is empty
func (q *Queue[T]) Front() (T, error) {
	var result T
	if q.head == nil {
		return result, ErrEmptyQueue
	} else {
		result = q.head.value
	}
	return result, nil
}

// Size returns the number of elements in the queue
func (q *Queue[T]) Size() int {

	return q.size
}

// IsEmpty returns true if the queue contains no elements
func (q *Queue[T]) IsEmpty() bool {
	return q.head == nil
}
