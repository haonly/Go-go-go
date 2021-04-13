package preprocess

import (
	"fmt"
	"io"
	"log"
	"strconv"
	"sync"
)

/*
[Problem 3]
전체 브랜드의 [transmission, fuel type] 조합의 tax 평균 구하기
*/

type taxSum struct {
	count int
	sum   int64
}

var TransmFuel map[string]*taxSum

func newTaxSum() *taxSum {
	newStruct := taxSum{}
	return &newStruct
}

func makeSet(wg *sync.WaitGroup, m *sync.Mutex, fileName string) {
	defer wg.Done()

	r := readFromCSV(fileName)
	// read one row of r
	count := 0
	for {
		record, err := r.Read()
		if count == 0 {
			count++
			continue
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error in r.Read(), err : %v\n", err)
		}
		key := fmt.Sprintf("%s_%s", record[TRANSMISSION], record[FUEL_TYPE])
		tax, err := strconv.Atoi(record[TAX])
		if err != nil {
			log.Fatalf("errror when strconv atoi tax\n")
		}

		m.Lock()

		// key가 처음인 경우 구조체 인스턴스 생성
		if TransmFuel[key] == nil {
			TransmFuel[key] = newTaxSum()
		}
		// key에 해당하는 sum, count 계산
		TransmFuel[key].sum += int64(tax)
		TransmFuel[key].count++

		m.Unlock()
	}
}

func calculateTaxAVG() int64 {
	var avg int64
	for _, v := range TransmFuel {
		// fmt.Printf("key : %v, tax AVG : %v\n", k, v.sum/int64(v.count))
		avg = v.sum / int64(v.count)
	}
	return avg
}
func Solution3Origin() {
	TransmFuel = make(map[string]*taxSum)
	var wg sync.WaitGroup
	var m sync.Mutex

	wg.Add(len(Brands))
	for _, v := range Brands {
		makeSet(&wg, &m, v)
	}
	wg.Wait()

	calculateTaxAVG()
	return
}

func Solution3GoRoutine() {
	TransmFuel = make(map[string]*taxSum)
	var wg sync.WaitGroup
	var m sync.Mutex

	wg.Add(len(Brands))
	for _, v := range Brands {
		go makeSet(&wg, &m, v)
	}
	wg.Wait()

	calculateTaxAVG()
	return
}
