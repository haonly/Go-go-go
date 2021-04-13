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
	FILE_DIR    = "../data/"
	FILE_FORMAT = "csv"
)

var Brands []string = []string{"audi", "bmw", "ford", "hyundi", "merc", "skoda", "toyota", "vauxhall", "vw"}

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
