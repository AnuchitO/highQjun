package benchmark

import "testing"

func BenchmarkCat(b *testing.B) {
	b.Log("BenchmarkCat")
}

func BenchmarkJoin(b *testing.B) {
	b.Log("BenchmarkJoin")
}
