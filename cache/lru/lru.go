package lru

import (
	"errors"
	"iter"

	"github.com/xyersh/xuyacs/list"
)

type Cache[K comparable, V any] interface {
	// Добавляет значение в кэш по ключу. Если ключ уже существует, обновляет значение.
	Put(key K, value V)

	// Возвращает значение по ключу и nil-ошибку, если ключ существует.
	// Если ключа нет, возвращает нулевое значение для V и error.
	Get(key K) (V, error)

	// Возвращает количество элементов в кэше
	Size() int

	// Чистит кэш
	Clear()

	// Возвращает итератор по элементам кэша
	All() iter.Seq2[K, V]
}

var (
	_              Cache[string, int] = (*CacheLRU[string, int])(nil)
	ErrKeyNodFound error              = errors.New("key not found")
)

type node[K comparable, V any] struct {
	key   K
	value V
}

type CacheLRU[K comparable, V any] struct {
	keyToElement map[K]*list.Element[*node[K, V]]
	capacity     int
	linkedList   *list.List[*node[K, V]]
}

func (c *CacheLRU[K, V]) GetList() *list.List[*node[K, V]] {
	return c.linkedList
}

func NewLRU[K comparable, V any](capacity int) *CacheLRU[K, V] {
	return &CacheLRU[K, V]{
		capacity:     capacity,
		linkedList:   list.New[*node[K, V]](),
		keyToElement: make(map[K]*list.Element[*node[K, V]], capacity),
	}
}

// All реализует интерфейс Cache
func (c *CacheLRU[K, V]) All() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {

		// итерируемся по списку от самой свежей записи до самой старой
		cur := c.linkedList.Front()
		for range c.Size() {
			n := c.getNodeFromElement(cur)
			if !yield(n.key, n.value) {
				return
			}
			cur = cur.Next()
		}
	}
}

// Clear реализует интерфейс Cache
func (c *CacheLRU[K, V]) Clear() {
	panic("unimplemented")
}

func (c *CacheLRU[K, V]) getNodeFromElement(element *list.Element[*node[K, V]]) *node[K, V] {
	return element.Value
}

// Get реализует интерфейс Cache.
func (c *CacheLRU[K, V]) Get(key K) (V, error) {
	if link, ok := c.keyToElement[key]; ok {

		//если ключ ЕСТЬ - переместим элемент в начало списка
		c.linkedList.MoveToFront(link)

		//если ключ ЕСТЬ - вернем значение + nil
		return c.getNodeFromElement(link).value, nil
	}

	//если ключа НЕТ - вернем nil + error
	var zero V
	return zero, ErrKeyNodFound

}

// Put реализует интерфейс Cache
func (c *CacheLRU[K, V]) Put(key K, value V) {
	// если ключ ЕСТЬ - обновим значение и переместим элемент в начало списка
	if link, ok := c.keyToElement[key]; ok {
		c.getNodeFromElement(link).value = value
		c.linkedList.MoveToFront(link)

		return
	}

	// если размер кэша уже равен capacity - удалим последний элемент
	var new_element *list.Element[*node[K, V]]
	if c.Size() == c.capacity {

		for_del := c.linkedList.Back()

		// удаляем элемент из мапы
		delete(c.keyToElement, c.getNodeFromElement(for_del).key)

		//переиспользуем существующий элемент и его ноду
		new_element = for_del
		new_element.Value.key = key
		new_element.Value.value = value

		// перемещаем элемент в начало списка
		c.linkedList.MoveToFront(new_element)
	} else {

		// создаем новую ноду, пинаем ее в начало списка
		nodeValue := &node[K, V]{key: key, value: value}
		new_element = c.linkedList.PushFront(nodeValue)
	}

	// добавляем элемент в мапу
	c.keyToElement[key] = new_element
}

// Size реализует интерфейс Cache
func (c *CacheLRU[K, V]) Size() int {
	return len(c.keyToElement)
}
