package bitarray

type BitArrayI interface {
	Set(idx int, val bool)
	Get(idx int) bool
}

type BitArray struct {
	data []byte
	size int
}

func NewBitArray(bitIdx int) *BitArray {
	return &BitArray{
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
