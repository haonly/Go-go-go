package main

import "fmt"

func main() {
	var some_number int
	fmt.Print("ENTER ANY NUMBER NOT 0 => ")
	fmt.Scanln(&some_number)

	fmt.Println("Num: ", some_number)
	switch {
	case some_number%3 == 0 && some_number%5 == 0:
		fmt.Println("FizzBuzzd")
	case some_number%3 == 0:
		fmt.Println("Fizz")
	case some_number%5 == 0:
		fmt.Println("Buzz")
	default:
		fmt.Println("Nothing")
	}

}
