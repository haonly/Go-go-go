package Solution12931

import (
	"strconv"
)

func Solution(n int) int{
	answer := 0

	for _, v := range strconv.Itoa(n){
		answer += int(v) - '0'
	}

	return answer
}
