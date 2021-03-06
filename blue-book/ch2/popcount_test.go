package main

import (
	"testing"

	"github.com/oaraujocesar/go-studies/blue-book/ch2/popcount"
)

// BenchmarkPopCountExpression-12 - 1000000000 - 0.2481 ns/op - 0 B/op - 0 allocs/op
func BenchmarkPopCountExpression(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountExpression(120)
	}
}

// BenchmarkPopCountLoop-12 -	206465565 - 6.010 ns/op - 0 B/op - 0 allocs/op
func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountLoop(120)
	}
}

// BenchmarkPopCountShiftValue-12 - 68561527 - 18.08 ns/op - 0 B/op - 0 allocs/op
func BenchmarkPopCountShiftValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountShiftValue(120)
	}
}

// BenchmarkPopCountClearRightMost-12 - 363935689 - 3.227 ns/op - 0 B/op - 0 allocs/op
func BenchmarkPopCountClearRightmost(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountClearRightmost(120)
	}
}
