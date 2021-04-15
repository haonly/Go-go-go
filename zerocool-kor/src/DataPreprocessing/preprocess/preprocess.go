package preprocess

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
)

const (
	MODEL = 1 * iota
	YEAR
	PRICE
	TRANSMISSION
	MILEAGE
	FUEL_TYPE
	TAX
	MPG
	ENGINE_SIZE
)

const (
	FILE_DIR    = "../data/"
	FILE_FORMAT = "csv"
)

var Brands []string = []string{"audi", "bmw", "ford", "hyundi", "merc", "skoda", "toyota", "vauxhall", "vw"}

// TODO : 1, 2, 3 번 솔루션 코드를 작성하고, 각 함수에서 답을 출력하세요.
// Preprocess, readFromCSV 함수는 예시입니다.바꾸셔도 좋아요.

func solution1() {
	log.Println("solution1 started")
	
	// read csv file of brands
	for _, brand := range Brands {
		total, avgPrice, avgMpg := GetAvgPriceMpgByBrand(brand)
		log.Printf("brand: %v, total: %v, price avg: %v, mpg avg: %v", brand, total, avgPrice, avgMpg)
	}
	log.Println("solution1 finished")
}

func solution1_goroutine() {
	log.Println("solution1_goroutine started")
	
	var wait sync.WaitGroup
	wait.Add(len(Brands))
	for _, brand := range Brands {
		go func(paramBrand string) {
			defer wait.Done()
			total, avgPrice, avgMpg := GetAvgPriceMpgByBrand(paramBrand)
			log.Printf("brand: %v, total: %v, price avg: %v, mpg avg: %v", paramBrand, total, avgPrice, avgMpg)
		}(brand)
	}
	wait.Wait()
	log.Println("solution1_goroutine finished")
}

func GetAvgPriceMpgByBrand(brand string) (int, int, float64) {
	log.Println("GetAvgPriceMpgByBrand brand:", brand)
	r := readFromCSV(brand)
	
	line := 0
	priceTotal := 0
	mpgTotal := 0.0
	
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}			
		if err != nil {
			log.Fatalf("error in r.Read(), err : %v\n", err)
		}
		line++
		if(line==1){
			continue
		}		

		mpg, _ := strconv.ParseFloat(record[MPG], 64)
		price, _ := strconv.Atoi(record[PRICE])
		mpgTotal += mpg
		priceTotal += price

		// log.Printf("record[%v] -> Model:%v, fuel_type:%v, mpg:%v, price:%v", line, record[MODEL], record[FUEL_TYPE], mpg, price)
	}
	line--
	return line, priceTotal/line, mpgTotal/float64(line)
}

func solution2() {}

func solution3() {}

func readFromCSV(fileName string) *csv.Reader {
	filePath := fmt.Sprint(FILE_DIR + fileName + "." + FILE_FORMAT)
	csvFile, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("error when open csv, err : %v\n", err)
	}
	r := csv.NewReader(csvFile)
	return r
}
