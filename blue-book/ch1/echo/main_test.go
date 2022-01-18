package main

import "testing"

// First result
// BenchmarkEcho1-12 - 63913 - 15660 ns/op - 424 B/op - 7 allocs/op
func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo1()
	}
}

// First result
// BenchmarkEcho2-12 - 60588 - 16576 ns/op - 424 B/op - 7 allocs/op
func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo2()
	}
}

// First result
// BenchmarkEcho3-12 - 78529 - 13950 ns/op - 152 B/op - 3 allocs/op
func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo3()
	}
}
