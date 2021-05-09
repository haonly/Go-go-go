package primefactor

func PrimeFactor(num int) []int {
	answer := []int{}
	numClone := num
	for i := 2; i < numClone; i++ {
		for num%i == 0 {
			num /= i
			answer = append(answer, i)
		}
	}
	return answer
}
