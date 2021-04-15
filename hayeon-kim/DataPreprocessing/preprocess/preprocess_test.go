package preprocess

import "testing"

func BenchmarkSolution1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution1()
	}
}

func BenchmarkSolution2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution2()
	}
}

func BenchmarkSolution3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution3()
	}
}