package main

import "testing"

func BenchmarkTest(b *testing.B) {
	matrix := createMatrix(100, 100, 3)

    for i := 0; i < b.N; i++ {
        find(matrix)
    }
}