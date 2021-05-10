package main

import "fmt"

func PrimeFactor(num int) []int {
	ret := make([]int, 0)
	for i := num; i > 1; i-- {
		if num%i == 0 {
			ret = append(ret, i)
		}
	}
	ret = append(ret, 1)
	return ret
}

func main() {
	var some_number int
	fmt.Print("ENTER ANY NUMBER => ")
	fmt.Scanln(&some_number)

	if some_number == 0 {
		fmt.Println("INPUT ZERO")
	} else {
		ret := PrimeFactor(some_number)
		fmt.Println(ret)
	}
}
