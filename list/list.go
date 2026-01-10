package list

import (
	"fmt"
	"iter"
	"strings"
)

// Element представляет узел в списке.
type Element[T any] struct {
	// Соседние элементы
	next, prev *Element[T]

	// Список, которому принадлежит этот элемент.
	list *List[T]

	// Значение, хранящееся в узле.
	Value T
}

// Next возвращает следующий элемент или nil.
func (e *Element[T]) Next() *Element[T] {
	if p := e.next; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

// Prev возвращает предыдущий элемент или nil.
func (e *Element[T]) Prev() *Element[T] {
	if p := e.prev; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

type ListI[T any] interface {
	Init() *List[T]
	Len() int
	Front() *Element[T]
	Back() *Element[T]
	PushFront(v T) *Element[T]
	PushBack(v T) *Element[T]
	Remove(e *Element[T]) T

	MoveToFront(e *Element[T])
	MoveToBack(e *Element[T])
	All() iter.Seq[T]
	String() string
}

// List представляет собой двухсвязный список.
// Нулевое значение List — это пустой список, готовый к использованию.
type List[T any] struct {
	root Element[T] // Страж (sentinel), упрощающий логику вставки/удаления
	len  int        // Текущая длина
}

// Init инициализирует или очищает список.
func (l *List[T]) Init() *List[T] {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

// lazyInit лениво инициализирует список, если он не был инициализирован.
func (l *List[T]) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

// New создает новый экземпляр списка.
func New[T any]() *List[T] {
	return (&List[T]{}).Init()
}

// Len возвращает количество элементов в списке.
func (l *List[T]) Len() int { return l.len }

// Front возвращает первый элемент списка или nil.
func (l *List[T]) Front() *Element[T] {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

// Back возвращает последний элемент списка или nil.
func (l *List[T]) Back() *Element[T] {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

// insert вставляет элемент e после at.
func (l *List[T]) insert(e, at *Element[T]) *Element[T] {
	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
	e.list = l
	l.len++
	return e
}

// insertValue — вспомогательный метод для вставки значения.
func (l *List[T]) insertValue(v T, at *Element[T]) *Element[T] {
	return l.insert(&Element[T]{Value: v}, at)
}

// PushFront вставляет значение в начало списка.
func (l *List[T]) PushFront(v T) *Element[T] {
	l.lazyInit()
	return l.insertValue(v, &l.root)
}

// PushBack вставляет значение в конец списка.
func (l *List[T]) PushBack(v T) *Element[T] {
	l.lazyInit()
	return l.insertValue(v, l.root.prev)
}

// Remove удаляет элемент из списка.
func (l *List[T]) Remove(e *Element[T]) T {
	if e.list == l {
		e.prev.next = e.next
		e.next.prev = e.prev
		e.next = nil // избегаем утечек памяти
		e.prev = nil
		e.list = nil
		l.len--
	}
	return e.Value
}

// move перемещает элемент e так, чтобы он оказался после элемента at.
// Элемент e уже должен находиться в списке.
func (l *List[T]) move(e, at *Element[T]) {
	if e == at {
		return
	}

	// Извлекаем e из текущей позиции
	e.prev.next = e.next
	e.next.prev = e.prev

	// Вставляем e после at
	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
}

// MoveToFront перемещает элемент e в начало списка.
// Если e не принадлежит списку l, список не изменяется.
func (l *List[T]) MoveToFront(e *Element[T]) {
	if e.list != l || l.root.next == e {
		return
	}
	// В списке со стражем (sentinel) начало — это сразу после root
	l.move(e, &l.root)
}

// MoveToBack перемещает элемент e в конец списка.
// Если e не принадлежит списку l, список не изменяется.
func (l *List[T]) MoveToBack(e *Element[T]) {
	if e.list != l || l.root.prev == e {
		return
	}
	// В списке со стражем конец — это перед root (его текущий prev)
	l.move(e, l.root.prev)
}

func (l *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := l.Front(); e != nil; e = e.Next() {
			if !yield(e.Value) {
				return
			}
		}
	}
}

func (l *List[T]) String() string {
	bldr := strings.Builder{}
	bldr.Write([]byte(fmt.Sprintf("type: %T   len: %d\n", l, l.Len())))
	for v := range l.All() {
		bldr.Write([]byte(fmt.Sprintf("%+v\n", v)))
	}

	return bldr.String()
}
