package preprocess

import "testing"

func BenchmarkSolution1(b *testing.B) {
	var brandNames = []string{"audi", "bmw", "ford", "hyundi", "merc", "skoda", "toyota", "vauxhall", "vw"}
	for i := 0; i < b.N; i++ {
		solution1(brandNames)
	}
}

func BenchmarkSolution2(b *testing.B) {
	var brandNames = []string{"audi", "bmw", "ford", "hyundi", "merc", "skoda", "toyota", "vauxhall", "vw"}
	for i := 0; i < b.N; i++ {
		solution2(brandNames)
	}
}

func BenchmarkSolution3(b *testing.B) {
	var brandNames = []string{"audi", "bmw", "ford", "hyundi", "merc", "skoda", "toyota", "vauxhall", "vw"}
	for i := 0; i < b.N; i++ {
		solution3(brandNames)
	}
}
