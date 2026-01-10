package bitarray

type BitArrayI interface {
	// Добавление бита по индексу
	Set(idx int, val bool)

	// Получение бита по индексу
	Get(idx int) bool
}

type BitArray struct {
	data []byte // биты храним здеся
	size uint   // размер массива
}

// Получаение нового биторого массива
func NewBitArray(bitCnt uint) *BitArray {
	return &BitArray{
		//в качестве длины слайса используется итоговое количество байт с учетом округления вверх
		data: make([]byte, (bitCnt+7)/8),
		size: bitCnt,
	}
}

func (b *BitArray) Set(bitIdx uint, val bool) {
	if bitIdx >= b.size {
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

func (b *BitArray) Get(bitIdx uint) bool {
	if bitIdx >= b.size {
		panic("index out of range")
	}

	byteIdx := bitIdx / 8
	bitOffset := uint(bitIdx % 8)

	return b.data[byteIdx]&(1<<bitOffset) != 0
}
