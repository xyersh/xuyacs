package bitarray

type BitArrayI interface {
	// Добавление бита по индексу
	Set(idx int, val bool)

	// Получение бита по индексу
	Get(idx int) bool
}

type BitArray struct {
	data []byte // биты храним здеся
	size int    // размер массива
}

// Получаение нового биторого массива
func NewBitArray(bitIdx int) *BitArray {
	return &BitArray{
		//в качестве длины слайса используется итоговое количество байт с учетом округления вверх
		data: make([]byte, (bitIdx+7)/8),
		size: bitIdx,
	}
}

func (b *BitArray) Set(bitIdx int, val bool) {
	if bitIdx < 0 || bitIdx >= b.size {
		panic("index out of range")
	}

	byteIdx := bitIdx / 8
	bitOffset := uint(bitIdx % 8)

	if val {
		// установить флаг
		b.data[byteIdx] |= 1 << (bitOffset)
	} else {
		// сбросить флаг
		b.data[byteIdx] &^= 1 << (bitOffset)
	}
}

func (b *BitArray) Get(bitIdx int) bool {
	if bitIdx < 0 || bitIdx >= b.size {
		panic("index out of range")
	}

	byteIdx := bitIdx / 8
	bitOffset := uint(bitIdx % 8)

	return b.data[byteIdx]&(1<<bitOffset) != 0
}
