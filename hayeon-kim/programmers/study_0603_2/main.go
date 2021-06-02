package study_0603_2

type QueueStruct struct {
	idx int
	priority int
}

type Queue []QueueStruct

func solution(priorities []int, location int) int {
	var ans = 0
	q := make(Queue, len(priorities))
	q = q.copyToQueue(priorities)


	for {
		idx, item := q[0].idx, q[0].priority
		q = q[1:]
		if len(q) >0 && getMax(q.getPriorityListFromQueue()) > item {
			q = q.queAppend(idx, item)

		} else{
			ans += 1
			if idx == location{
				break
			}

		}
	}
	return ans
}


func (q Queue) copyToQueue(priorities []int) Queue {
	for idx, val := range priorities {
		q[idx].idx = idx
		q[idx].priority = val
	}
	return q
}


func getMax(inputList []int) int {
	var max = 0
	for _, val := range inputList {
		if val > max {
			max = val
		}
	}
	return max
}


func (q Queue) getPriorityListFromQueue() []int {
	var list []int
	for _, val := range q {
		list = append(list, val.priority)
	}
	return list
}


func (q Queue) queAppend(idx int, item int) Queue{
	appendItem := QueueStruct{idx:idx, priority:item}
	q = append(q, appendItem)
	return q
}


