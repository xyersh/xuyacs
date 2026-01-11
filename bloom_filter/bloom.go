package bloom_filter

import (
	"math"

	"github.com/spaolacci/murmur3"
	"github.com/xyersh/xuyacs/bitarray"
)

type BloomFilterI interface {
	Add([]byte)       // добавить отпечаток знаяения в фильтр Блума
	Test([]byte) bool // проверить наличие отпечатка в фильтре Блума (возможен ложно-положительный результат )
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
// p - вероятность ложного срабатывания ( от нуля  до единицы)
func NewBloomFilter(n int, p float64) *BloomFilter {
	if p < 0.0 || p > 1.0 {
		panic("`p` must be between 0.0 and 1.0")
	}

	m := uint(math.Ceil(-float64(n) * math.Log(p) / math.Pow(math.Log(2), 2)))
	k := uint(math.Ceil(float64(m) / float64(n) * math.Log(2)))

	if k == 0 {
		k = 1
	}

	bitsArr := make(Bitmap, k)
	for i := 0; i < int(k); i++ {
		bitsArr[i] = *bitarray.NewBitArray(m)
	}

	return &BloomFilter{
		bitmap: bitsArr,
		m:      m,
		k:      k,
	}

}

func (bf *BloomFilter) hash128(data []byte) (uint, uint) {
	h1, h2 := murmur3.Sum128(data)
	return uint(h1), uint(h2)
}

func (bf *BloomFilter) Add(data []byte) {
	h1, h2 := bf.hash128(data)

	for i := 0; i < int(bf.k); i++ {
		index := (h1 + uint(i)*h2) % bf.m

		bf.bitmap[i].Set(index, true)
	}
}

func (bf *BloomFilter) Test(data []byte) bool {
	h1, h2 := bf.hash128(data)

	for i := 0; i < int(bf.k); i++ {
		index := (h1 + uint(i)*h2) % bf.m

		if !bf.bitmap[i].Get(index) {
			return false
		}
	}

	return true
}
