// https://programmers.co.kr/learn/courses/30/lessons/42586

package Solution42586

func Solution(progresses []int, speeds []int) []int {

	answer := []int{}

	productionDay := 0   // 현 시점의 배포일
	productionCount := 0 // productionDay에 배포하는 개수

	for i := 0; i < len(progresses); i++ {
		days := (100 - progresses[i]) / speeds[i]
		lastDay := (100 - progresses[i]) % speeds[i]
		if lastDay > 0 {
			// 진행이 남아, 하루 더 추가되는 경우
			days += 1
		}

		if productionDay == 0 {
			// 최초 실행되는 경우
			productionDay = days
			productionCount += 1
			continue
		} else if productionDay < days {
			// 현재 배포일 이후에 배포되는 경우
			answer = append(answer, productionCount) // 새로운 production 추가
			productionCount = 1                      // count 초기화
			productionDay = days
		} else {
			// 현재 배포일 이전, 포함 배포되는 경우
			productionCount += 1
		}
	}

	// 추가해야 할 배포 개수가 남아있는 경우
	if productionCount != 0 {
		answer = append(answer, productionCount)
	}

	return answer
}

// (100 - 93) / 1 = 7
// (100 - 93) % 1 = 0
// 7 + 0 = 7 days

// (100 - 30) = 70
// 70 / 30 = 2
// 70 % 30 = 10 (!0)
// 2 + 1 = 3 days

// (100 - 55) = 45
// 45 / 5 = 9
// 45 % 5 = 0
// 9 days

// [7, 3, 9] -> [2, 1]
