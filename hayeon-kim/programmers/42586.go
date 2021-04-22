func solution(progresses []int, speeds []int) []int {
    depDayList := []int{}
    answer := []int{}

    for i := 0; i < len(progresses); i++{
        days := (100-progresses[i]) / speeds[i]
        mods := (100-progresses[i]) % speeds[i]
        if mods != 0{
            days += 1
        }
        depDayList = append(depDayList, days)
    }
    front := 0
    for idx := 0; idx < len(progresses); idx ++ {
        if depDayList[front] < depDayList[idx]{
            answer = append(answer, idx-front)
            front = idx
        }
    }
    answer = append(answer, len(depDayList) - front)

    return answer
}