package circle_buffer

import (
	"slices"
	"testing"
)

func TestNew(t *testing.T) {
	cb := New[int](5)
	tests := []struct {
		name string
		want any
	}{
		{"capasity", 5},
		{"length", 0},
		{"is empty", true},
		{"is full", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "capasity" {
				if cb.Cap() != tt.want {
					t.Errorf("Expected capacity %d, got %d", tt.want, cb.Cap())
				}
			}
			if tt.name == "length" {
				if cb.Len() != tt.want {
					t.Errorf("Expected length %d, got %d", tt.want, cb.Len())
				}
			}
			if tt.name == "is empty" {
				if cb.IsEmpty() != tt.want {
					t.Errorf("Expected is empty %t, got %t", tt.want, cb.IsEmpty())
				}
			}
			if tt.name == "is full" {
				if cb.IsFull() != tt.want {
					t.Errorf("Expected is full %t, got %t", tt.want, cb.IsFull())
				}
			}
		})
	}
}

func TestEnqueueDequeue(t *testing.T) {
	cb := New[int](3)

	// Push items
	err := cb.Enqueue(1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	err = cb.Enqueue(2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	err = cb.Enqueue(3)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if cb.Len() != 3 {
		t.Errorf("Expected length 3, got %d", cb.Len())
	}
	if !cb.IsFull() {
		t.Error("Buffer should be full")
	}

	// Pop items in FIFO order
	val, err := cb.Dequeue()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val != 1 {
		t.Errorf("Expected 1, got %d", val)
	}

	val, err = cb.Dequeue()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val != 2 {
		t.Errorf("Expected 2, got %d", val)
	}

	val, err = cb.Dequeue()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val != 3 {
		t.Errorf("Expected 3, got %d", val)
	}

	if !cb.IsEmpty() {
		t.Error("Buffer should be empty after popping all")
	}
}

func TestEnqueueFullBuffer(t *testing.T) {
	cb := New[int](2)

	cb.Enqueue(1)
	cb.Enqueue(2)

	err := cb.Enqueue(3)
	if err == nil {
		t.Error("Expected error when pushing to full buffer")
	}
	if !cb.IsFull() {
		t.Error("Buffer should be full")
	}
}

func TestPopEmptyBuffer(t *testing.T) {
	cb := New[int](2)

	_, err := cb.Dequeue()
	if err == nil {
		t.Error("Expected error when popping from empty buffer")
	}
	if !cb.IsEmpty() {
		t.Error("Buffer should be empty")
	}
}

func TestWrapAround(t *testing.T) {
	cb := New[int](3)

	// Fill buffer
	cb.Enqueue(1)
	cb.Enqueue(2)
	cb.Enqueue(3)

	// Pop one to make space
	val, _ := cb.Dequeue()
	if val != 1 {
		t.Errorf("Expected 1, got %d", val)
	}

	// Push one more to test wrap around
	err := cb.Enqueue(4)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Pop remaining
	val, _ = cb.Dequeue()
	if val != 2 {
		t.Errorf("Expected 2, got %d", val)
	}
	val, _ = cb.Dequeue()
	if val != 3 {
		t.Errorf("Expected 3, got %d", val)
	}
	val, _ = cb.Dequeue()
	if val != 4 {
		t.Errorf("Expected 4, got %d", val)
	}
}

func TestLenCap(t *testing.T) {
	cb := New[string](4)

	if cb.Cap() != 4 {
		t.Errorf("Expected capacity 4, got %d", cb.Cap())
	}
	if cb.Len() != 0 {
		t.Errorf("Expected length 0, got %d", cb.Len())
	}

	cb.Enqueue("a")
	cb.Enqueue("b")
	if cb.Len() != 2 {
		t.Errorf("Expected length 2, got %d", cb.Len())
	}

	cb.Dequeue()
	if cb.Len() != 1 {
		t.Errorf("Expected length 1, got %d", cb.Len())
	}
}

func TestIsEmptyIsFull(t *testing.T) {
	cb := New[string](2)

	if !cb.IsEmpty() {
		t.Error("New buffer should be empty")
	}
	if cb.IsFull() {
		t.Error("New buffer should not be full")
	}

	cb.Enqueue("x")
	if cb.IsEmpty() {
		t.Error("Buffer should not be empty after push")
	}
	if cb.IsFull() {
		t.Error("Buffer should not be full yet")
	}

	cb.Enqueue("y")
	if cb.IsEmpty() {
		t.Error("Buffer should not be empty")
	}
	if !cb.IsFull() {
		t.Error("Buffer should be full")
	}

	cb.Dequeue()
	if cb.IsEmpty() {
		t.Error("Buffer should not be empty after pop")
	}
	if cb.IsFull() {
		t.Error("Buffer should not be full after pop")
	}

	cb.Dequeue()
	if !cb.IsEmpty() {
		t.Error("Buffer should be empty after popping all")
	}
	if cb.IsFull() {
		t.Error("Buffer should not be full")
	}
}

// Test with different types to ensure generics work
func TestGenerics(t *testing.T) {
	cb := New[[]int](2)

	slice1 := []int{1, 2}
	slice2 := []int{3, 4, 5}

	cb.Enqueue(slice1)
	cb.Enqueue(slice2)

	pop1, _ := cb.Dequeue()
	if !slices.Equal(pop1, slice1) {
		t.Errorf("Expected %v, got %v", slice1, pop1)
	}

	pop2, _ := cb.Dequeue()
	if !slices.Equal(pop2, slice2) {
		t.Errorf("Expected %v, got %v", slice2, pop2)
	}
}
