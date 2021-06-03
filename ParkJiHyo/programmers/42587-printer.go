type Queue struct {
	Items chan int
	count int
}

func (q *Queue) push(item int) {
	q.Items <- item
	q.count++
}
func (q *Queue) front() int {
	if q.count == 0 {
		return -1
	}
	q.count--
	return <-q.Items
}

func solution(priorities []int, location int) int {
	var answer int = 0

	picked := make([]int, 0)
	queue := Queue{
		Items: make(chan int, len(priorities)),
		count: 0,
	}
	for i := 0; i < len(priorities); i++ {
		queue.push(i)
	}
	for {

		curIndex := queue.front()
		if curIndex == -1 {
			break
		}
		if priorities[curIndex] != max_element(priorities) {
			queue.push(curIndex)
		} else {
			picked = append(picked, curIndex)
			priorities[curIndex] = 0
		}
	}
	for i := 0; i < len(priorities); i++ {
		if picked[i] == location {
			answer = i + 1
			break
		}
	}
	return answer
}
func max_element(arr []int) int {
	var result int = -987654321
	for i := 0; i < len(arr); i++ {
		if result < arr[i] {
			result = arr[i]
		}
	}
	return result
}