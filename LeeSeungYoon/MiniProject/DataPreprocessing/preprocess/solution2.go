package preprocess

import (
	"fmt"
	"io"
	"log"
	"strconv"
	"sync"
)

/*
[Problem 2]
전체 브랜드에 대해서,
transmission이 Manual, Automatic, Semi-Auto 인 차량 각각의 mileage, mpg의 평균 구하기
*/

type transmCh struct {
	mpgCH  chan float64
	milgCh chan int
}

type sumVal struct {
	milgSUM int64
	mpgSUM  int64
	count   int
}

var transmissionSUM map[string]*sumVal
var manualCh *transmCh
var autoCh *transmCh
var semiAutoCH *transmCh

func newTransmCh() *transmCh {
	ch := transmCh{}
	ch.milgCh = make(chan int, 100000)
	ch.mpgCH = make(chan float64, 100000)
	return &ch
}

func closeTransmCh(ch *transmCh) {
	close(ch.milgCh)
	close(ch.mpgCH)
}

func inputChannel(wg *sync.WaitGroup, m *sync.Mutex, fileName string) {
	defer wg.Done()

	r := readFromCSV(fileName)
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

		mileage, err := strconv.Atoi(record[MILEAGE])
		if err != nil {
			log.Fatalf("error strconv.Atoi record[MILEAGE], err : %v\n", err)
		}
		mpg, err := strconv.ParseFloat(record[MPG], 64)
		if err != nil {
			log.Fatalf("error strconv.Atoi record[MPG], err : %v\n", err)
		}

		m.Lock() // acquire lock

		// Transmission 칼럼값에 따라 각 채널로 mileage, mpg 값 저장
		switch record[TRANSMISSION] {
		case "Manual":
			manualCh.milgCh <- mileage
			manualCh.mpgCH <- mpg
		case "Automatic":
			autoCh.milgCh <- mileage
			autoCh.mpgCH <- mpg
		case "Semi-Auto":
			semiAutoCH.milgCh <- mileage
			semiAutoCH.mpgCH <- mpg
			// default:
			// log.Printf("record[TRANSMISSION] is not among [manual, auto, semi-auto], TRANSMISSION : %v\n", TRANSMISSION)
		}

		m.Unlock() // release lock
	}
}

func calculateSum(wg *sync.WaitGroup, ch *transmCh, transmissionType string) {
	defer func() {
		wg.Done()
	}()

	var transType string = fmt.Sprintf("%sSUM", transmissionType) // manualSUM, autoSUM, semiAutoSUM
	for v := range ch.milgCh {
		transmissionSUM[transType].milgSUM += int64(v)
		transmissionSUM[transType].count++
	}
	for v := range ch.mpgCH {
		transmissionSUM[transType].mpgSUM += int64(v)
	}
}

func caculateSumAVG() (int64, int64, int64, int64, int64, int64) {
	manualMilgAVG := transmissionSUM["manualSUM"].milgSUM / int64(transmissionSUM["manualSUM"].count)
	manualMpgAVG := transmissionSUM["manualSUM"].mpgSUM / int64(transmissionSUM["manualSUM"].count)

	autoMilgAVG := transmissionSUM["autoSUM"].milgSUM / int64(transmissionSUM["autoSUM"].count)
	autoMpgAVG := transmissionSUM["autoSUM"].mpgSUM / int64(transmissionSUM["autoSUM"].count)

	semiAutoMilgAVG := transmissionSUM["semiAutoSUM"].milgSUM / int64(transmissionSUM["semiAutoSUM"].count)
	semiAutoMpgAVG := transmissionSUM["semiAutoSUM"].mpgSUM / int64(transmissionSUM["semiAutoSUM"].count)

	// log.Printf("avg manual mileage : %v, mpg : %v\n", transmissionSUM["manualSUM"].milgSUM/int64(transmissionSUM["manualSUM"].count), transmissionSUM["manualSUM"].mpgSUM/int64(transmissionSUM["manualSUM"].count))
	// log.Printf("avg auto mileage : %v, mpg : %v\n", transmissionSUM["autoSUM"].milgSUM/int64(transmissionSUM["autoSUM"].count), transmissionSUM["autoSUM"].mpgSUM/int64(transmissionSUM["autoSUM"].count))
	// log.Printf("avg semiAuto mileage : %v, mpg : %v\n", transmissionSUM["semiAutoSUM"].milgSUM/int64(transmissionSUM["semiAutoSUM"].count), transmissionSUM["semiAutoSUM"].mpgSUM/int64(transmissionSUM["semiAutoSUM"].count))
	return manualMilgAVG, manualMpgAVG, autoMilgAVG, autoMpgAVG, semiAutoMilgAVG, semiAutoMpgAVG
}

func Solution2GoRoutine() {

	manualCh = newTransmCh()
	autoCh = newTransmCh()
	semiAutoCH = newTransmCh()

	// map 인스턴스 생성 및 구조체 초기화
	transmissionSUM = make(map[string]*sumVal)
	transmissionSUM["manualSUM"] = &sumVal{}
	transmissionSUM["autoSUM"] = &sumVal{}
	transmissionSUM["semiAutoSUM"] = &sumVal{}

	var wg sync.WaitGroup
	var m sync.Mutex

	// 각 brand 데이터 -> 채널에 저장
	wg.Add(len(Brands))
	for _, brand := range Brands {
		go inputChannel(&wg, &m, brand)
	}
	wg.Wait()

	// 채널 닫기
	closeTransmCh(manualCh)
	closeTransmCh(autoCh)
	closeTransmCh(semiAutoCH)

	// 저장된 채널로부터 Sum 구하기
	wg.Add(3)
	go calculateSum(&wg, manualCh, "manual")
	go calculateSum(&wg, autoCh, "auto")
	go calculateSum(&wg, semiAutoCH, "semiAuto")
	wg.Wait()

	// 평균값 구하기
	caculateSumAVG()

	return
}

func Solution2Origin() {

	manualCh = newTransmCh()
	autoCh = newTransmCh()
	semiAutoCH = newTransmCh()

	transmissionSUM = make(map[string]*sumVal)
	transmissionSUM["manualSUM"] = &sumVal{}
	transmissionSUM["autoSUM"] = &sumVal{}
	transmissionSUM["semiAutoSUM"] = &sumVal{}

	var wg sync.WaitGroup
	var m sync.Mutex

	wg.Add(len(Brands))
	for _, brand := range Brands {
		inputChannel(&wg, &m, brand)
	}
	// wg.Wait()

	closeTransmCh(manualCh)
	closeTransmCh(autoCh)
	closeTransmCh(semiAutoCH)

	wg.Add(3)
	calculateSum(&wg, manualCh, "manual")
	calculateSum(&wg, autoCh, "auto")
	calculateSum(&wg, semiAutoCH, "semiAuto")
	// wg.Wait()

	caculateSumAVG()

	return
}

// 1. 3개 종류 transmission에 대한 채널 만들기 (각각 mpg, milg)
// 2. 브랜드별 csv 파일을 파싱해서 채널에 값 넣기
// 3. 파싱이 모두 끝나면, 채널에 쌓인 값들로 평균 계산하기
// manual
// automatic
// semi auto
