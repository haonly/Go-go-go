package preprocess

import "testing"

func BenchmarkSolution1Origin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solution1Origin()
	}
}

func BenchmarkSolution1GoRoutine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solution1GoRoutine()
	}
}

func BenchmarkSolution2Origin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solution2Origin()
	}
}

func BenchmarkSolution2GoRoutine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solution2GoRoutine()
	}
}

func BenchmarkSolution3Origin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solution3Origin()
	}
}
func BenchmarkSolution3GoRoutine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solution3GoRoutine()
	}
}
