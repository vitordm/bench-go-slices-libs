package main

import "testing"

var slicesOfNumbers []int = []int{1, 4, 5, 6, 7, 9, 23, 55, 293}

func BenchmarkContainsNumberWithFunk(b *testing.B) {
	println("Starting... ", b.N)
	for i := 0; i < b.N; i++ {
		ContainsNumberWithFunk(slicesOfNumbers, i)
	}
}

func BenchmarkContainsNumberWithSlice(b *testing.B) {
	println("Starting... ", b.N)
	for i := 0; i < b.N; i++ {
		ContainsNumberWithSlice(slicesOfNumbers, i)
	}
}
