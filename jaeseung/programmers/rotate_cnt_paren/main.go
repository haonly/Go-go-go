package main

import (
	"log"
	"math/rand"
	"sync"
)

func solution(s string) int {
	ans := 0
	N := len(s)
	result := make(chan int, N)
	wg := sync.WaitGroup{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		str := rotate(i, s)
		go func() {
			defer wg.Done()
			checkValidPair(result, str)
		}()
	}
	go func() {
		wg.Wait()
		close(result)
	}()
	for ret := range result {
		ans += ret
	}
	return ans
}

func rotate(i int, s string) string {
	if i == 0 {
		return s
	}
	if i == len(s)-1 {
		return s[i:i+1] + s[:i]
	}
	return s[i:] + s[:i]
}

type Stack struct {
	top  int
	elem []string
}

func NewStack() *Stack {
	return &Stack{top: -1}
}

func (s *Stack) IsEmpty() bool {
	return s.top == -1
}

func (s *Stack) Pop() string {
	ret := s.elem[s.top]
	s.elem = s.elem[:s.top]
	s.top--
	return ret
}

func (s *Stack) Push(str string) {
	s.elem = append(s.elem, str)
	s.top++
}

func checkValidPair(ret chan<- int, s string) {
	stack := NewStack()
	for _, each := range s {
		each := string(each)
		switch each {
		case "(":
			stack.Push(each)
		case "{":
			stack.Push(each)
		case "[":
			stack.Push(each)
		case ")":
			if stack.IsEmpty() {
				ret <- 0
				return
			}
			val := stack.Pop()
			if val != "(" {
				ret <- 0
				return
			}
		case "}":
			if stack.IsEmpty() {
				ret <- 0
				return
			}
			val := stack.Pop()
			if val != "{" {
				ret <- 0
				return
			}
		case "]":
			if stack.IsEmpty() {
				ret <- 0
				return
			}
			val := stack.Pop()
			if val != "[" {
				ret <- 0
				return
			}
		}
	}
	if !stack.IsEmpty() {
		ret <- 0
		return
	}

	ret <- 1
}

func main() {
	s := NewStack()
	N := 5000
	testInput := ""
	openings := []string{"(", "{", "["}
	for i := 0; i < N; i++ {
		v := openings[rand.Int()%3]
		testInput += v
		s.Push(v)
	}

	for !s.IsEmpty() {
		each := s.Pop()
		switch each {
		case "(":
			testInput += ")"
		case "{":
			testInput += "}"
		case "[":
			testInput += "]"
		}
	}
	log.Println(testInput)
}
