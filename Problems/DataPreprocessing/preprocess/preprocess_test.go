package preprocess

import "testing"

func BenchmarkSolution1(b *testing.B) {
	var brandNames = []string{"audi", "bmw", "ford", "hyundi", "merc", "skoda", "toyota", "vauxhall", "vw"}
	const desiredNumTask = 10
	for i := 0; i < b.N; i++ {
		solution1(brandNames, desiredNumTask)
	}
}

func BenchmarkSolution2(b *testing.B) {
	var brandNames = []string{"audi", "bmw", "ford", "hyundi", "merc", "skoda", "toyota", "vauxhall", "vw"}
	for i := 0; i < b.N; i++ {
		solution2(brandNames) // Test 대상함수, 예) 카프카 메시지 처리 함수
	}
}

func BenchmarkSolution3(b *testing.B) {
	var brandNames = []string{"audi", "bmw", "ford", "hyundi", "merc", "skoda", "toyota", "vauxhall", "vw"}
	for i := 0; i < b.N; i++ {
		solution3(brandNames)
	}
}
