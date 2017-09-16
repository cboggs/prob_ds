package main

import "fmt"

func stringToInt(s string) uint64 {
	x := uint64(0)
	for _, r := range s {
		x += uint64(r)
	}
	return x
}

func hashCountOddFromRight(x uint64, n uint64) uint64 {
	y := uint64(0)

	for i := 0; i < 16; i++ {
		oddBit := x & 1
		y |= (oddBit << uint64(i))
		x >>= 2
	}

	return y % n
}

func hashCountEvenFromRight(x uint64, n uint64) uint64 {
	y := uint64(0)

	for i := 0; i < 16; i++ {
		x >>= 1
		evenBit := x & 1
		y |= (evenBit << uint64(i))
		x >>= 1
	}

	return y % n
}

func insertString(s string, e uint64, n uint64) ([]uint64, error) {
	i := stringToInt(s)
	h1 := hashCountOddFromRight(i, n)
	h2 := hashCountEvenFromRight(i, n)

	fmt.Printf("empty: %064b\n", e)
	//fmt.Printf("1>>h1: %64b\n", 1>>h1)
	e |= (uint64(1) << h1)
	fmt.Printf("addh1: %064b\n", e)

	//fmt.Printf("1>>h2: %64b\n", 1>>h2)
	e |= (uint64(1) << h2)
	fmt.Printf("addh2: %064b\n", e)

	return []uint64{h1, h2}, nil
}

func main() {
	n := uint64(64)
	e := uint64(0)
	a, err := insertString("hello world", e, n)
	if err != nil {
		fmt.Printf("returnedErr: %+v\n", err)
	}
	fmt.Printf("a: %+v\n", a)
}
