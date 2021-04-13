package preprocess

import (
	"fmt"
	"log"
	"sync"
)

type extransmCh struct {
	mpgCH  chan float64
	milgCh chan int
}

type exStruct struct {
	n int64
}

var dict map[string]*exStruct

func exnewTransmCh() *extransmCh {
	ch := extransmCh{}
	ch.milgCh = make(chan int)
	ch.mpgCH = make(chan float64)
	return &ch
}

func inputCh(ch *extransmCh, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		ch.milgCh <- i
	}
	close(ch.milgCh)
	wg.Done()
}

type chStruct struct {
	ch1 chan int
	ch2 chan int
}

func makeStruct() *chStruct {
	ch := chStruct{}
	ch.ch1 = make(chan int, 10)
	ch.ch2 = make(chan int, 10)
	return &ch
}

var exChannel *chStruct

func callee() {
	exChannel.ch1 <- 1
	exChannel.ch2 <- 2
	fmt.Println("<-ch1 : ", <-exChannel.ch1)
	fmt.Println("<-ch2 : ", <-exChannel.ch2)
}

func Example1() {
	exChannel = makeStruct()
	callee()
}

func Example2() {
	log.Println("example start...")

	var wg sync.WaitGroup
	ch := exnewTransmCh()

	wg.Add(1)
	go inputCh(ch, &wg)

	go func() {
		for v := range ch.milgCh {
			fmt.Printf("%v ", v)
		}
		fmt.Println()
	}()

	// for v := range ch.milgCh {
	// 	fmt.Printf("%v ", v)
	// }
	// fmt.Println()
	// wg.Wait()

	fmt.Println("____________________")
	dict = make(map[string]*exStruct)
	dict["key1"] = &exStruct{}
	dict["key2"] = &exStruct{}
	dict["key3"] = &exStruct{}

	dict["key1"].n += 1

	for k, v := range dict {
		log.Printf("key : %v, value : %v\n", k, v)
	}

	log.Println("example end...")
}

func Example3() {
	log.Println("example 3 start....")
	var wg sync.WaitGroup
	ch := make(chan int, 20)

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go ex3Callee(&wg, ch, i)
	}
	wg.Wait()

	wg.Add(1)
	close(ch)
	go func() {
		for v := range ch {
			fmt.Printf("%v ", v)
		}
		defer wg.Done()
	}()

	wg.Wait()
	fmt.Println()
	log.Println("example 3 end....")
}

func ex3Callee(wg *sync.WaitGroup, ch chan int, i int) {
	ch <- i
	defer wg.Done()
}
