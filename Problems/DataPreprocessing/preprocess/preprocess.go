package preprocess

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

const (
	MODEL = iota
	YEAR
	PRICE
	TRANSMISSION
	MILEAGE
	FUEL_TYPE
	TAX
	MPG
	ENGINE_SIZE
)

// TODO : 1, 2, 3 번 솔루션 코드를 작성하고, 각 함수에서 답을 출력하세요.
// Preprocess, readFromCSV 함수는 예시입니다.바꾸셔도 좋아요.

func solution1(brandNames []string) {
	wg := sync.WaitGroup{}
	for _, brandName := range brandNames {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			r := readFromCSV(name)
			all, err := r.ReadAll()
			if err != nil {
				log.Fatal(err)
				return
			}
			recordCnt := len(all) - 1
			totalPrice := 0
			var totalMpg float64

			for _, record := range all[1:] {
				price, _ := strconv.Atoi(record[PRICE])
				totalPrice += price
				mpg, _ := strconv.ParseFloat(record[MPG], 64)
				totalMpg += mpg
			}
			priceAvg := float64(totalPrice) / float64(recordCnt)
			mpgAvg := totalMpg / float64(recordCnt)
			log.Println(name, " -> price AVG : ", priceAvg, ", mpg AVG : ", mpgAvg)
		}(brandName)
	}
	wg.Wait()
}

func solution2(brandNames []string) {
	wg := sync.WaitGroup{}
	for _, brandName := range brandNames {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			transMap := make(map[string]bool)
			cntMap := make(map[string]int)
			milMap := make(map[string]int)
			mpgMap := make(map[string]float64)
			r := readFromCSV(name)
			all, err := r.ReadAll()
			if err != nil {
				log.Fatal(err)
				return
			}
			for _, record := range all[1:] {
				transType := record[TRANSMISSION]
				transMap[transType] = true
				cntMap[transType] += 1
				mil, _ := strconv.Atoi(record[MILEAGE])
				milMap[transType] += mil
				mpg, _ := strconv.ParseFloat(record[MPG], 64)
				mpgMap[transType] += mpg
			}

			for key := range transMap {
				milAvg := float64(milMap[key]) / float64(cntMap[key])
				mpgAvg := mpgMap[key] / float64(cntMap[key])
				log.Println(name, " ", key, " -> mileage AVG : ", milAvg, ", mpg AVG : ", mpgAvg)
			}
		}(brandName)
	}
	wg.Wait()
}

func solution3(brandNames []string) {
	wg := sync.WaitGroup{}
	for _, brandName := range brandNames {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			cntPerTransAndFuelType := make(map[string]int)
			taxPerTransAndFuelType := make(map[string]int)
			r := readFromCSV(name)
			all, err := r.ReadAll()
			if err != nil {
				log.Fatal(err)
				return
			}
			for _, record := range all[1:] {
				k := record[TRANSMISSION] + "," + record[FUEL_TYPE]
				cntPerTransAndFuelType[k] += 1
				tax, _ := strconv.Atoi(record[TAX])
				taxPerTransAndFuelType[k] += tax
			}

			for key := range taxPerTransAndFuelType {
				taxAvg := float64(taxPerTransAndFuelType[key]) / float64(cntPerTransAndFuelType[key])
				log.Println(name, " ", key, " -> tax AVG : ", taxAvg)
			}
		}(brandName)
	}
	wg.Wait()
}

func Preprocess(brandNames []string) {
	solution1(brandNames)
	solution2(brandNames)
	solution3(brandNames)
}

func readFromCSV(fileName string) *csv.Reader {
	const (
		//FILE_DIR    = "./Problems/DataPreprocessing/data/"
		FILE_DIR = "../data/"
		//FILE_DIR    = "./data/"
		FILE_FORMAT = "csv"
	)
	filePath := fmt.Sprint(FILE_DIR + fileName + "." + FILE_FORMAT)
	csvFile, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("error when open csv, err : %v\n", err)
	}
	r := csv.NewReader(csvFile)
	return r
}
