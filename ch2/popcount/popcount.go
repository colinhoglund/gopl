package main

import (
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// LoopPopCount returns the population count (number of set bits) of x.
func LoopPopCount(x uint64) int {
	var count byte
	for _, i := range [8]byte{0, 1, 2, 3, 4, 5, 6, 7} {
		count += pc[byte(x>>(i*8))]
	}
	return int(count)
}

// ShiftPopCount returns the population count (number of set bits) of x.
func ShiftPopCount(x uint64) int {
	count := x & 1
	for x > 0 {
		x = x >> 1
		count += x & 1
	}
	return int(count)
}

// ClearPopCount returns the population count (number of set bits) of x.
func ClearPopCount(x uint64) int {
	var count int
	for x > 0 {
		count++
		x = x & (x - 1)
	}
	return count
}

func main() {
	fmt.Println(PopCount(11111111))
	fmt.Println(LoopPopCount(11111111))
	fmt.Println(ShiftPopCount(11111111))
	fmt.Println(ClearPopCount(11111111))
}
