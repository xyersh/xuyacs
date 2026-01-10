package main

import (
	"fmt"

	"github.com/xyersh/xuyacs/bitarray"
)

func main() {
	ba := bitarray.NewBitArray(64)

	ba.Set(1, true)
	ba.Set(3, true)
	ba.Set(4, true)
	fmt.Printf("%+v\n", ba)

	fmt.Printf("[0]: %v\n", ba.Get(0))
	fmt.Printf("[1]: %v\n", ba.Get(1))
	fmt.Printf("[2]: %v\n", ba.Get(2))
	fmt.Printf("[3]: %v\n", ba.Get(3))
	fmt.Printf("[4]: %v\n", ba.Get(4))
}
