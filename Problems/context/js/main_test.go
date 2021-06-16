package main

import (
	"context"
	"github.com/stretchr/testify/suite"
	"log"
	"sync"
	"testing"
	"time"
)

func TestAaa(t *testing.T) {
	suite.Run(t, new(AaaTestSuite))
}

type AaaTestSuite struct {
	suite.Suite
}

func (s *AaaTestSuite) SetupTest() {

}

func (s *AaaTestSuite) TearDownTest() {

}

func (s *AaaTestSuite) TestA() {
	log.Println("이 요청 3초 이내에 끝나야해")
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(3*time.Second))
	defer cancel()

	resultChan := make(chan string, 1)

	go func(ctx context.Context, result chan<- string) {
		log.Println("DB작업 ~~ 1초 뒤 끝~~")
		time.Sleep(1 * time.Second)
		result <- "끝"
	}(ctx, resultChan)

	var wasOk bool
	select {
	case ret := <-resultChan:
		log.Printf("DB작업 끝? %s", ret)
		wasOk = true
	case <-ctx.Done():
		log.Printf("context error=%v", ctx.Err())
		log.Println("DB작업이 제때 안 끝났네,, 이번 작업은 실패라고 알려야지...")
		wasOk = false
	}

	if wasOk {
		log.Println("DB가 시간안에 잘 끝났어")
	} else {
		log.Println("DB 관련 동작이 오래 걸렸나봐, context가 timeout나서 이번 요청 실패야")
	}

	// 성공 경우(DB작업 함수 1초 대기)
	// 2021/06/10 16:24:10 이 요청 3초 이내에 끝나야해
	// 2021/06/10 16:24:10 DB작업 ~~ 1초 뒤 끝~~
	// 2021/06/10 16:24:12 DB작업 끝? 끝
	// 2021/06/10 16:24:12 DB가 시간안에 잘 끝났어

	// 실패 경우(DB작업 함수 4초 대기)
	// 2021/06/10 16:23:40 이 요청 3초 이내에 끝나야해
	// 2021/06/10 16:23:40 DB작업 ~~ 4초 뒤 끝~~
	//	2021/06/10 16:23:43 context error=context deadline exceeded
	// 2021/06/10 16:23:43 DB작업이 제때 안 끝났네,, 이번 작업은 실패라고 알려야지...
	// 2021/06/10 16:23:43 DB 관련 동작이 오래 걸렸나봐, context가 timeout나서 이번 요청 실패야

	// 결과
	// 2021/06/10 15:57:54 Waiting 3 sec
	// 2021/06/10 15:57:56 Waiting 5 sec, Cancel called
	// 2021/06/10 15:57:56 context canceled
	wg := sync.WaitGroup{}
	wg.Add(2)
	ctx = context.WithValue(context.Background(), "request-id", "req-id-123")

	go func(ctx context.Context) {
		defer wg.Done()
		log.Println("Func-A")
		log.Println(ctx.Value("request-id"))
	}(ctx)
	go func(ctx context.Context) {
		defer wg.Done()
		log.Println("Func-B")
		log.Println(ctx.Value("request-id"))
	}(ctx)
	wg.Wait()

	// 2021/06/10 16:29:56 Func-B
	// 2021/06/10 16:29:56 req-id-123
	// 2021/06/10 16:29:56 Func-A
	// 2021/06/10 16:29:56 req-id-123
}
