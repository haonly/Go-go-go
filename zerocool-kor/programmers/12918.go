package main

import (
	"log"
	"strconv"
)

func solution(s string) bool {
	log.Println("solution s:",s)
	if(len(s)>=1 && len(s)<=8){
		_, err := strconv.Atoi(s)
		return !(err!=nil)
	}else{
		return false
	}
}

func main() {
	log.Println("result -> ", solution("a234"))
	log.Println("result -> ", solution("1452"))
	log.Println("result -> ", solution("0000"))
}