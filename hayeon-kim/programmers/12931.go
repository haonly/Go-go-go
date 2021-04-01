import (
    "strconv"
)

func solution(n int) int {
    answer := 0
    
    var strN string = strconv.Itoa(n)
    for i := 0; i < len(strN); i++ {
        val, _ := strconv.Atoi(string(strN[i]))
        answer += val
    }
//     for {
//         answer += n % 10
//         n /= 10

//         if n == 0 {
//             break
//         }
//     }

    return answer
}