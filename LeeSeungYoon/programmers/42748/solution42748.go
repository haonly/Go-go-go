// https://programmers.co.kr/learn/courses/30/lessons/42748

package Solution42748

func partition(array *[]int, low, high int) int {
	var pivot int = (*array)[high]
	var i int = low - 1

	for j := low; j < high; j++ {
		if (*array)[j] < pivot {
			i++
			(*array)[i], (*array)[j] = (*array)[j], (*array)[i]
		}
	}
	(*array)[i+1], (*array)[high] = (*array)[high], (*array)[i+1]
	return i + 1
}

func quickSort(array *[]int, low, high int) {
	if low < high {
		var pivotIndex int = partition(array, low, high)
		quickSort(array, low, pivotIndex-1)
		quickSort(array, pivotIndex+1, high)
	}
}

func Solution(array []int, commands [][]int) []int {
	answer := []int{}

	for i := range commands {
		start := commands[i][0] - 1
		end := commands[i][1]
		k := commands[i][2]

		// array 슬라이싱
		slicedArray := make([]int, end-start)
		copy(slicedArray, array[start:end])

		// 퀵 정렬
		quickSort(&slicedArray, 0, len(slicedArray)-1)

		// k번째 수 구하기
		answer = append(answer, slicedArray[k-1])
	}
	return answer
}
