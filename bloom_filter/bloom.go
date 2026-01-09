package bloom_filter

import (
	"math"
	// "github.com/spaolacci/murmur3"
	"github.com/xyersh/xuyacs/bitarray"
)

type BloomFilterI interface {
	Add([]byte)
	Test([]byte) bool
}

type Bitmap []bitarray.BitArray

type BloomFilter struct {
	bitmap Bitmap
	// n      float64 	// ожидаемое количество элементов для анализа
	m uint // размер битового массива (в битах)
	k uint // количество битовых массивов в bitmap
}

// NewBloomFilter создает новый фильтр Блума с заданным количеством элементов и вероятностью ложного срабатывания.
// n - ожидаемое количество элементов для анализа
// p - вероятность ложного срабатывания (доя от единицы)
func NewBloomFilter(n int, p float64) *BloomFilter {
	m := uint(math.Ceil(-float64(n) * math.Log(p) / math.Pow(math.Log(2), 2)))
	k := uint(math.Ceil(float64(m) / float64(n) * math.Log(2)))

	if k == 0 {
		k = 1
	}

	bitsArr := make(Bitmap, k)
	for i := 0; i < int(k); i++ {
		bitsArr[i] = *bitarray.NewBitArray(int(m))
	}

	return &BloomFilter{
		bitmap: bitsArr,
		m:      m,
		k:      k,
	}

}

func (bf *BloomFilter) Add(data []byte) {
}

func (bf *BloomFilter) Test(data []byte) bool {
	return false
}
