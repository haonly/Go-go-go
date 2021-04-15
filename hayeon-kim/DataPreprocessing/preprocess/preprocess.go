package preprocess

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
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
	FILE_DIR    = "data/"
	FILE_FORMAT = "csv"
)

var Brands []string = []string{"audi", "bmw", "cclass", "focus", "ford", "hyundi", "merc", "skoda", "toyota", "vauxhall", "vw"}

// TODO : 1, 2, 3 번 솔루션 코드를 작성하고, 각 함수에서 답을 출력하세요.
// Preprocess, readFromCSV 함수는 예시입니다.바꾸셔도 좋아요.

func solution1() {}

func solution2() {}

func solution3() {}

func Preprocess() {
	// read csv file of brands
	for _, fileName := range Brands {
		// csv reader
		r := readFromCSV(fileName)
		log.Println("brand : ", fileName)

		// read one row of r
		for {
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("error in r.Read(), err : %v\n", err)
			}
			log.Printf("record[%v] : %v\n", MODEL, record[MODEL])
			log.Printf("record[%v] : %v\n", FUEL_TYPE, record[FUEL_TYPE])

		}
		break
	}
	return
}

func readFromCSV(fileName string) *csv.Reader {
	filePath := fmt.Sprint(FILE_DIR + fileName + "." + FILE_FORMAT)
	csvFile, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("error when open csv, err : %v\n", err)
	}
	r := csv.NewReader(csvFile)
	return r
}