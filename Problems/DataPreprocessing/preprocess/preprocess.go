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

type priceAndMpg struct {
	price int
	mpg   float64
}

func solution1(brandNames []string, numTasks int) {
	wg := sync.WaitGroup{}
	for _, brandName := range brandNames {
		wg.Add(1)
		go calcPriceAndMPGAVG(&wg, brandName, numTasks)
	}
	wg.Wait()
}

func calcPriceAndMPGAVG(wg *sync.WaitGroup, name string, desiredNumTask int) {
	defer wg.Done()
	taskCh := make(chan priceAndMpg)

	all := readRecords(name)

	subWg, recordCnt := buildSyncGroup(desiredNumTask, all)

	divideWorkload(recordCnt, desiredNumTask, subWg, taskCh, all)

	// Fan in result
	totalPrice, totalMpg := collectResult(taskCh)

	// Result
	printResult(totalPrice, recordCnt, totalMpg, name)
}

func readRecords(name string) [][]string {
	r := readFromCSV(name)
	all, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return all
}

func buildSyncGroup(numTasks int, all [][]string) (sync.WaitGroup, int) {
	subWg := sync.WaitGroup{}
	subWg.Add(numTasks)
	recordCnt := len(all) - 1
	if recordCnt%numTasks != 0 {
		subWg.Add(1)
	}
	return subWg, recordCnt
}

func divideWorkload(recordCnt int, numTasks int, subWg sync.WaitGroup, taskCh chan priceAndMpg, all [][]string) {
	start := 1
	offset := recordCnt / numTasks
	end := start + offset
	for {
		go func(start, end int) {
			defer subWg.Done()
			calcSumForPriceAndMPG(taskCh, all, start, end)
		}(start, end)

		if end == recordCnt+1 {
			break
		}

		start += offset
		end += offset

		if end > recordCnt {
			end = recordCnt + 1
		}
	}
	go func() {
		subWg.Wait()
		close(taskCh)
	}()
}

func collectResult(taskCh chan priceAndMpg) (int, float64) {
	var totalPrice int
	var totalMpg float64
	for result := range taskCh {
		totalPrice += result.price
		totalMpg += result.mpg
	}
	return totalPrice, totalMpg
}

func printResult(totalPrice, recordCnt int, totalMpg float64, name string) {
	priceAvg := float64(totalPrice) / float64(recordCnt)
	mpgAvg := totalMpg / float64(recordCnt)
	log.Println(name, " -> price AVG : ", priceAvg, ", mpg AVG : ", mpgAvg)
}

func calcSumForPriceAndMPG(c chan priceAndMpg, records [][]string, start, end int) {
	totalPrice := 0
	var totalMpg float64
	for _, record := range records[start:end] {
		price, _ := strconv.Atoi(record[PRICE])
		totalPrice += price
		mpg, _ := strconv.ParseFloat(record[MPG], 64)
		totalMpg += mpg
	}
	sumPriceAndMpg := priceAndMpg{
		price: totalPrice,
		mpg:   totalMpg,
	}
	c <- sumPriceAndMpg
}

func solution2(brandNames []string) {
	wg := sync.WaitGroup{}
	for _, brandName := range brandNames {
		wg.Add(1)
		go calcMileageAndMPGAverage(&wg, brandName)
	}
	wg.Wait()
}

func calcMileageAndMPGAverage(wg *sync.WaitGroup, name string) {
	defer wg.Done()
	transMap := make(map[string]bool)
	cntMap := make(map[string]int)
	milMap := make(map[string]int)
	mpgMap := make(map[string]float64)
	records := readRecords(name)
	_calcMileageAndMPGAverage(records, milMap, mpgMap, transMap, cntMap)
	printMileageAndMPGResult(milMap, mpgMap, transMap, cntMap, name)
}

func _calcMileageAndMPGAverage(records [][]string, milMap map[string]int, mpgMap map[string]float64, transMap map[string]bool, cntMap map[string]int) {
	for _, record := range records[1:] {
		transType := record[TRANSMISSION]
		transMap[transType] = true
		cntMap[transType] += 1
		mil, _ := strconv.Atoi(record[MILEAGE])
		milMap[transType] += mil
		mpg, _ := strconv.ParseFloat(record[MPG], 64)
		mpgMap[transType] += mpg
	}
}

func printMileageAndMPGResult(milMap map[string]int, mpgMap map[string]float64, transMap map[string]bool, cntMap map[string]int, name string) {
	for key := range transMap {
		milAvg := float64(milMap[key]) / float64(cntMap[key])
		mpgAvg := mpgMap[key] / float64(cntMap[key])
		log.Println(name, " ", key, " -> mileage AVG : ", milAvg, ", mpg AVG : ", mpgAvg)
	}
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
	//const numTasks = 10
	//solution1(brandNames, numTasks)
	//solution2(brandNames)
	//solution3(brandNames)
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
