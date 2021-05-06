package main

import (
	"context"
	"log"
	"math/rand"
	"sync"
	"time"
)

func solQ1(wg *sync.WaitGroup) {
	defer wg.Done()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	go reqUser(ctx)
	go reqInventoryWithRandomFail(cancel, ctx)
	result := make(chan int)
	go reqSlow(ctx, result)

	select {
	case <-result:
		log.Println("Request succeed")
		ctx.Done()
	case <-ctx.Done():
		log.Printf("case <-timeoutCtx.Done()=%v\n", ctx.Err())
	}
}

func reqUser(ctx context.Context) {
	defer log.Println("REQ-user done...")
	defer ctx.Done()
	time.Sleep(200 * time.Millisecond)
}

func reqInventoryWithRandomFail(cancel context.CancelFunc, ctx context.Context) {
	defer log.Println("REQ-inventory done...")
	defer ctx.Done()
	time.Sleep(500 * time.Millisecond)
	rand.Seed(int64(time.Now().Second()))
	if rand.Intn(2)%2 == 0 {
		defer log.Println("REQ-inventory failed...")
		cancel()
	}
}

func reqSlow(ctx context.Context, c chan int) {
	defer log.Println("REQ-slow done...")
	defer ctx.Done()
	time.Sleep(time.Second)
	c <- 1
}

func solQ2(wg *sync.WaitGroup) {
	defer wg.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 800*time.Millisecond)
	defer cancel()

	go reqUser(ctx)
	go reqInventory(ctx)
	result := make(chan int)
	go reqSlowWithRandomFail(ctx, result)

	select {
	case <-result:
		log.Println("Okkkk")
		ctx.Done()
	case <-ctx.Done():
		log.Printf("case <-timeoutCtx.Done()=%v\n", ctx.Err())
	}
}

func reqInventory(ctx context.Context) {
	defer log.Println("REQ-inventory done...")
	defer ctx.Done()
	time.Sleep(500 * time.Millisecond)
}

func reqSlowWithRandomFail(ctx context.Context, ch chan int) {
	defer log.Println("REQ-slow done...")
	defer ctx.Done()
	rand.Seed(int64(time.Now().Second()))
	if rand.Intn(3)%2 == 0 {
		log.Println("REQ-slow too slow...")
		time.Sleep(time.Second)
	}
	time.Sleep(500 * time.Millisecond)
	ch <- 1
}

func withTimeCheck(q func(wg *sync.WaitGroup)) {
	start := time.Now()
	wg := sync.WaitGroup{}
	wg.Add(1)
	q(&wg)
	wg.Wait()
	elapsed := time.Since(start)
	log.Println(elapsed)
}

func main() {
	log.Println("Question#1")
	withTimeCheck(solQ1)
	log.Println("Question#2")
	withTimeCheck(solQ2)
}
