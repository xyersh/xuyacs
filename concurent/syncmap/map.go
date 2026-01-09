package syncmap

import (
	"sync"
	"unsafe"
)

type MapShard[K comparable, V any] struct {
	items map[K]V
	sync.RWMutex
}

type SyncMap[K comparable, V any] struct {
	shardCnt int               // степень двойки, полученное число  из которой определяет количество шардов
	shards   []*MapShard[K, V] // шарды - здеся
}

func NewSyncMap[K comparable, V any](shardCnt int) *SyncMap[K, V] {
	shards := make([]*MapShard[K, V], shardCnt)
	for i := 0; i < shardCnt; i++ {
		shards[i] = &MapShard[K, V]{items: make(map[K]V)}
	}
	return &SyncMap[K, V]{shardCnt: shardCnt, shards: shards}
}

func (sm *SyncMap[K, V]) fnv32(key K) uint32 {
	// Определяем размер типа T в памяти
	size := unsafe.Sizeof(key)

	// Получаем указатель на данные ключа и представляем их как срез байтов
	// Внимание: это работает для простых типов и структур без указателей внутри
	var data []byte
	data = unsafe.Slice((*byte)(unsafe.Pointer(&key)), size)

	hash := uint32(2166136261)
	const prime32 = 16777619
	for _, b := range data {
		hash *= prime32
		hash ^= uint32(b)
	}
	return hash
}

func (sm *SyncMap[K, V]) GetShard(key K) *MapShard[K, V] {
	hash := sm.fnv32(key)
	return sm.shards[hash&(uint32(sm.shardCnt)-1)] // оптимизация. A&(B-1) == A%B, если B степень двойки
}

func (sm *SyncMap[K, V]) Set(key K, value V) {
	shard := sm.GetShard(key)
	shard.Lock()
	shard.items[key] = value
	shard.Unlock()
}

// Get получает значение по ключу
func (sm *SyncMap[K, V]) Get(key K) (V, bool) {
	shard := sm.GetShard(key)
	shard.RLock()
	val, ok := shard.items[key]
	shard.RUnlock()
	return val, ok
}
