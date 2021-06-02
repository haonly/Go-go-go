// https://programmers.co.kr/learn/courses/30/lessons/70128

package Solution42587

// 1. 초기 index = 0
// 2 .현재 index 보다 큰 max 수 존재하면, 날리기
// + 날린 수는 0으로 만들고, count + 1
// 3. 날린 수의 index와 location 값이 같으면, count 리턴

func findBiggerMax(pivot int, priorities []int) (bool, int) {

	isExist := false
	maxIndex := pivot
	for i := 0; i < len(priorities); i++ {
		if i == pivot || priorities[i] == 0 {
			continue
		}
		if priorities[i] > priorities[maxIndex] {
			maxIndex = i
			isExist = true
		}
	}
	return isExist, maxIndex
}

func makeMaxZero(index int, priorities *[]int) {
	ZERO := 0
	(*priorities)[index] = ZERO
}

func Solution(priorities []int, location int) int {

	count := 0
	pivot := 0

	for {

		// log.Printf("pivot : %v, count : %v, priority : %v\n", pivot, count, priorities)

		isExist, maxIndex := findBiggerMax(pivot, priorities)
		count += 1

		if isExist {
			pivot = maxIndex
		}
		// max index와 찾던 location 값이 같다면
		if pivot == location {
			break
		}

		// log.Printf("make deletion index : %v\n", pivot)
		makeMaxZero(pivot, &priorities)

		pivot += 1
		if pivot == len(priorities) {
			pivot = 0
		}
	}
	return count
}
