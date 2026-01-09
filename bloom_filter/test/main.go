package main

import (
	"fmt"

	"github.com/xyersh/xuyacs/bloom_filter"
)

func main() {
	bf := bloom_filter.NewBloomFilter(1000000, 5.0)
	fmt.Println(bf)

}
