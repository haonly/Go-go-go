package main

import (
	"log"
)

type ParenthesesStack []int8

func (s *ParenthesesStack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *ParenthesesStack) Push(ele int8) {
	*s = append(*s, ele)
}

func (s *ParenthesesStack) Pop() (int8, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

var openParentheses = map[string]bool{
	"(": true,
	"{": true,
	"[": true,
}
var pmap = map[string]int8{
	"(": 1,
	"{": 2,
	"[": 3,
	")": -1,
	"}": -2,
	"]": -3,
}

func isValid(s string) bool {
	log.Println("isValid(", s, ") started")
	var openList ParenthesesStack
	for _, ch := range s {
		letter := string(ch)
		//log.Println("i:",i,", letter:", letter)
		n, found := pmap[letter]
		if !found {
			return false
		}
		_, open := openParentheses[letter]
		if open {
			openList.Push(n)
		} else {
			poped, _ := openList.Pop()
			if poped+n != 0 {
				return false
			}
		}
	}
	return true
}

func main() {
	log.Println("result -> ", isValid("abcd"))
	log.Println("result -> ", isValid("()"))
	log.Println("result -> ", isValid("(){}[]({[]})"))

}
