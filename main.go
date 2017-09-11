package main

import (
	"fmt"
)

func hashCountOddFromRight(x uint8, n uint8) uint8 {
	y := uint8(0)

	for i := range []uint8{0, 2, 4, 6} {
		oddBit := x & 1
		y |= (oddBit << uint8(i))
		x >>= 2
	}

	return y % n
}

func main() {
	x := uint8(247)
	n := uint8(11)

	fmt.Printf("x>>0: %08b\n", x)
	h1 := hashCountOddFromRight(x, n)
	fmt.Printf("h1: %d\n", h1)
}
