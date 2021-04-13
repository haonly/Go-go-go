package preprocess

import (
	"io"
	"log"
	"strconv"
)

/*
[Problem 1]
각 브랜드별 price, mpg 평균 구하기
*/

func Solution1Origin() {

	for _, branch := range Brands {
		// priceAVG, mpgAVG := eachBrandAVG(branch)
		eachBrandAVG(branch)
		// log.Printf("priceAVG : %v, mpgAVG : %v\n", priceAVG, mpgAVG)
	}
	return
}

func Solution1GoRoutine() {

	for _, branch := range Brands {
		// priceAVG, mpgAVG := eachBrandAVG(branch)
		go eachBrandAVG(branch)
		// log.Printf("priceAVG : %v, mpgAVG : %v\n", priceAVG, mpgAVG)
	}
	return
}
func eachBrandAVG(fileName string) (int, int) {

	priceSum := 0
	mpgSum := 0.0

	r := readFromCSV(fileName)

	// read one row of r
	count := 0
	for {
		record, err := r.Read()
		if count == 0 {
			count++
			continue
		}
		count++

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error in r.Read(), err : %v\n", err)
		}
		price, _ := strconv.Atoi(record[PRICE])
		priceSum += price

		mpg, _ := strconv.ParseFloat(record[MPG], 64)
		mpgSum += mpg

	}
	return priceSum / count, int(mpgSum) / count
}
