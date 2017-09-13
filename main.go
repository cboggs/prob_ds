package main

import (
	"fmt"
)

func stringToInt(s string) uint32 {
	x := uint32(0)
	for _, r := range s {
		x += uint32(r)
	}
	return x
}

func hashCountOddFromRight(x uint32, n uint32) uint32 {
	y := uint32(0)

	for i := 0; i < 16; i++ {
		oddBit := x & 1
		y |= (oddBit << uint32(i))
		x >>= 2
	}

	return y % n
}

func hashCountEvenFromRight(x uint32, n uint32) uint32 {
	y := uint32(0)

	for i := 0; i < 16; i++ {
		x >>= 1
		evenBit := x & 1
		y |= (evenBit << uint32(i))
		x >>= 1
	}

	return y % n
}

func main() {
	s := "hello world"
	n := uint32(32)

	sInt := stringToInt(s)
	fmt.Printf("sInt: %d\n", sInt)
	fmt.Printf("bin(sInt): %08b\n", sInt)
	h1 := hashCountOddFromRight(sInt, n)
	h2 := hashCountEvenFromRight(sInt, n)
	fmt.Printf("h1: %d\n", h1)
	fmt.Printf("h2: %d\n", h2)
}
