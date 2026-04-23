package circle_buffer

import "fmt"

type CircleBuffer[T any] interface {
	Push(T) error
	Pop() (T, error)
	Len() int
	Cap() int
	IsEmpty() bool
	IsFull() bool
}

type circleBuffer[T any] struct {
	buffer []T
	head   int
	tail   int
	len    int
	cap    int
}

func New[T any](cap int) *circleBuffer[T] {
	buffer := make([]T, cap)
	return &circleBuffer[T]{
		buffer: buffer,
		head:   0,
		tail:   0,
		len:    0,
		cap:    cap,
	}
}

func (c *circleBuffer[T]) Push(item T) error {
	if c.len == c.cap {
		return fmt.Errorf("buffer is full")
	}
	c.buffer[c.tail] = item
	c.tail = (c.tail + 1) % c.cap
	c.len++
	return nil
}

func (c *circleBuffer[T]) Pop() (T, error) {
	var zero T
	if c.len == 0 {
		return zero, fmt.Errorf("buffer is empty")
	}
	item := c.buffer[c.head]
	c.head = (c.head + 1) % c.cap
	c.len--
	return item, nil
}

func (c *circleBuffer[T]) Len() int {
	return c.len
}

func (c *circleBuffer[T]) Cap() int {
	return c.cap
}

func (c *circleBuffer[T]) IsEmpty() bool {
	return c.len == 0
}

func (c *circleBuffer[T]) IsFull() bool {
	return c.len == c.cap
}
