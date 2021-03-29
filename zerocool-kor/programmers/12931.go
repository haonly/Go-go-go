package main

import (
	"log"
	"strconv"
)

func solution(n int) int {
	log.Println("solution n:",n)
	sum:=0
	for _, c := range(strconv.Itoa(n)) {
		n,_ := strconv.Atoi(string(c))
		sum += n
	}
	return sum
}

func main() {
	log.Println("result -> ", solution(123))
	log.Println("result -> ", solution(987))
}