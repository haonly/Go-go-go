package FizzBuzz

import "fmt"

func FizBuz(num int) string {

	answer := ""

	var mod3 int = num % 3
	var mod5 int = num % 5

	if mod3 == 0 && mod5 == 0 {
		answer = "FizzBuzz"
	} else if mod3 == 0 {
		answer = "Fizz"
	} else if mod5 == 0 {
		answer = "Buzz"
	}
	fmt.Println(answer)
	return answer
}
