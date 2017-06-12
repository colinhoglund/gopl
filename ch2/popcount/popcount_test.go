package main

import (
	"testing"
)

// BenchmarkPopCount evaluates function performance
func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(11111111)
	}
}

// BenchmarkLoopPopCount evaluates function performance
func BenchmarkLoopPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LoopPopCount(11111111)
	}
}

// BenchmarkShiftPopCount evaluates function performance
func BenchmarkShiftPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ShiftPopCount(11111111)
	}
}

// BenchmarkClearPopCount evaluates function performance
func BenchmarkClearPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ClearPopCount(11111111)
	}
}
