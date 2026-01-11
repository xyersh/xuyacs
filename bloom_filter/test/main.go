package main

import (
	"fmt"

	"github.com/xyersh/xuyacs/bloom_filter"
)

func main() {
	bf := bloom_filter.NewBloomFilter(100, 0.5)

	bf.Add([]byte("Sasha"))
	bf.Add([]byte("Masha"))
	bf.Add([]byte("Dasha"))
	bf.Add([]byte("Natasha"))

	fmt.Printf("Sasha in filter: %t\n", bf.Test([]byte("Sasha")))
	fmt.Printf("Masha in filter: %t\n", bf.Test([]byte("Masha")))
	fmt.Printf("Dasha in filter: %t\n", bf.Test([]byte("Dasha")))
	fmt.Printf("Natasha in filter: %t\n", bf.Test([]byte("Natasha")))

	fmt.Printf("Lesha in filter: %t\n", bf.Test([]byte("Lesha")))
	fmt.Printf("Kesha in filter: %t\n", bf.Test([]byte("Kesha")))
}
